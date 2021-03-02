package circuit_breaker

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"
)

// ServiceBreaker：熔断器结构
type ServiceBreaker struct {
	mu               sync.RWMutex // mu：读写锁
	name             string       // name：熔断器名字
	state            State        // 熔断器状态
	windowInterval   time.Duration
	metrics          Metrics
	tripStrategyFunc TripStrategyFunc
	halfMaxCalls     uint64
	stateOpenTime    time.Time
	sleepTimeout     time.Duration
	stateChangeHook  func(name string, fromState State, toState State)
}

type State int

const (
	StateClosed State = iota
	StateOpen
	StateHalfOpen
)

func (s State) String() string {
	switch s {
	case StateClosed:
		return "closed"
	case StateHalfOpen:
		return "half-open"
	case StateOpen:
		return "open"
	default:
		return fmt.Sprintf("unknown state:%d", s)
	}
}

// NewServiceBreaker：初始化熔断器
func NewServiceBreaker(op Option) (*ServiceBreaker, error) {
	if op.WindowInterval <= 0 || op.HalfMaxCalls <= 0 || op.SleepTimeout <= 0 {
		return nil, errors.New("incomplete options")
	}
	breaker := new(ServiceBreaker)
	breaker.name = op.Name
	breaker.windowInterval = op.WindowInterval
	breaker.halfMaxCalls = op.HalfMaxCalls
	breaker.sleepTimeout = op.SleepTimeout
	breaker.stateChangeHook = op.StateChangeHook
	breaker.tripStrategyFunc = ChooseTrip(&op.TripStrategy)
	breaker.nextWindow(time.Now())
	return breaker, nil
}

var (
	ErrStateOpen    = errors.New("service breaker is open")
	ErrTooManyCalls = errors.New("service breaker is halfopen, too many calls")
)

func (breaker *ServiceBreaker) beforeCall() error {
	breaker.mu.Lock()
	defer breaker.mu.Unlock()
	now := time.Now()
	switch breaker.state {
	case StateOpen:
		//after sleep timeout, can retry
		if breaker.stateOpenTime.Add(breaker.sleepTimeout).Before(now) {
			log.Printf("%s 熔断过冷却期，尝试半开", breaker.name)
			breaker.changeState(StateHalfOpen, now)
			return nil
		}
		log.Printf("%s 熔断打开，请求被阻止", breaker.name)
		return ErrStateOpen
	case StateHalfOpen:
		if breaker.metrics.CountAll >= breaker.halfMaxCalls {
			log.Printf("%s 熔断半开，请求过多被阻止", breaker.name)
			return ErrTooManyCalls
		}
	default: //Closed
		if !breaker.metrics.WindowTimeStart.IsZero() && breaker.metrics.WindowTimeStart.Before(now) {
			breaker.nextWindow(now)
			return nil
		}
	}
	return nil
}

func (breaker *ServiceBreaker) Call(exec func() (interface{}, error)) (interface{}, error) {
	//before call
	err := breaker.beforeCall()
	if err != nil {
		return nil, err
	}
	//if panic occur
	defer func() {
		err := recover()
		if err != nil {
			breaker.afterCall(false)
			panic(err)
		}
	}()
	//call
	breaker.metrics.OnCall()
	result, err := exec()
	//after call
	breaker.afterCall(err == nil)
	return result, err
}

func (breaker *ServiceBreaker) afterCall(success bool) {
	breaker.mu.Lock()
	defer breaker.mu.Unlock()
	if success {
		breaker.onSuccess(time.Now())
	} else {
		breaker.onFail(time.Now())
	}
}

func (breaker *ServiceBreaker) onSuccess(now time.Time) {
	breaker.metrics.OnSuccess()
	if breaker.state == StateHalfOpen && breaker.metrics.ConsecutiveSuccess >= breaker.halfMaxCalls {
		breaker.changeState(StateClosed, now)
	}
}

func (breaker *ServiceBreaker) changeState(state State, now time.Time) {
	if breaker.state == state {
		return
	}
	prevState := breaker.state
	breaker.state = state
	//goto next window,reset metrics
	breaker.nextWindow(time.Now())
	//record open time
	if state == StateOpen {
		breaker.stateOpenTime = now
	}
	//callback hook
	if breaker.stateChangeHook != nil {
		breaker.stateChangeHook(breaker.name, prevState, state)
	}
}

func (breaker *ServiceBreaker) onFail(now time.Time) {
	breaker.metrics.OnFail()
	switch breaker.state {
	case StateClosed:
		if breaker.tripStrategyFunc(breaker.metrics) {
			breaker.changeState(StateOpen, now)
		}
	case StateHalfOpen:
		breaker.changeState(StateOpen, now)

	}
}

func (breaker *ServiceBreaker) nextWindow(now time.Time) {
	breaker.metrics.NewBatch()
	breaker.metrics.OnReset() //clear count num
	var zero time.Time
	switch breaker.state {
	case StateClosed:
		if breaker.windowInterval == 0 {
			breaker.metrics.WindowTimeStart = zero
		} else {
			breaker.metrics.WindowTimeStart = now.Add(breaker.windowInterval)
		}
	case StateOpen:
		breaker.metrics.WindowTimeStart = now.Add(breaker.sleepTimeout)
	default: //halfopen
		breaker.metrics.WindowTimeStart = zero //halfopen no window
	}
}
