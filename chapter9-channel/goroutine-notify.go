package main

import (
	"fmt"
	"math/rand"
	"runtime"
)

func GenerateIntA(done chan struct{}) chan int {
	ch := make(chan int)
	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}
		}
		// 收到通知
		close(ch)
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntA(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 发送通知
	close(done)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	// 生产者已退出
	println("NumGoroutine=", runtime.NumGoroutine())
}
