package main

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
)

// 定时channel操作
// ToChanTimedContext 定时channel操作【context 优先】
func ToChanTimedContext(ctx context.Context, d time.Duration, message reflect.Type, c chan<- reflect.Type) (written bool) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	select {
	case c <- message:
		return true
	case <-ctx.Done():
		return false
	}
}

// ToChanTimedTimer 定时channel操作【timer】
func ToChanTimedTimer(d time.Duration, message reflect.Type, c chan<- reflect.Type) (written bool) {
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case c <- message:
		return true
	case <-t.C:
		return false
	}
}

// 先进先出服务
// FirstComeFirstServedSelect 先进先出【select 优先】
func FirstComeFirstServedSelect(message reflect.Type, a, b chan<- reflect.Type) {
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		}
	}
}

// FirstComeFirstServedGoroutines 先进先出【goroutines 优先】
func FirstComeFirstServedGoroutines(message reflect.Type, a, b chan<- reflect.Type) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { a <- message; wg.Done() }()
	go func() { b <- message; wg.Done() }()
	wg.Wait()
}

// FirstComeFirstServedGoroutinesVariadic 先进先出【不清楚需要开启多少goroutines】
func FirstComeFirstServedGoroutinesVariadic(message reflect.Type, chs ...chan<- reflect.Type) {
	var wg sync.WaitGroup
	wg.Add(len(chs))
	for _, c := range chs {
		c := c
		go func() { c <- message; wg.Done() }()
	}
	wg.Wait()
}

// 整合到一起
// ToChansTimedTimerSelect 【整合到一起 timer+select，明确知道数量】
func ToChansTimedTimerSelect(d time.Duration, message reflect.Type, a, b chan reflect.Type) (written int) {
	t := time.NewTimer(d)
	for i := 0; i < 2; i++ {
		select {
		case a <- message:
			a = nil
		case b <- message:
			b = nil
		case <-t.C:
			return i
		}
	}
	t.Stop()
	return 2
}

// ToChansTimedContextGoroutines 【整合到一起 context+goroutine,不清楚chan数量】
func ToChansTimedContextGoroutines(ctx context.Context, d time.Duration, message reflect.Type, ch ...chan reflect.Type) (written int) {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()
	var (
		wr int32
		wg sync.WaitGroup
	)
	wg.Add(len(ch))
	for _, c := range ch {
		c := c
		go func() {
			defer wg.Done()
			select {
			case c <- message:
				atomic.AddInt32(&wr, 1)
			case <-ctx.Done():
			}
		}()
	}
	wg.Wait()
	return int(wr)
}
