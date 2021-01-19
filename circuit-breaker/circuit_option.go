package circuit_breaker

import "time"

type Option struct {
	Name            string
	WindowInterval  time.Duration
	HalfMaxCalls    uint64
	SleepTimeout    time.Duration
	StateChangeHook func(name string, fromState State, toState State)
	TripStrategy    TripStrategyOption
}
