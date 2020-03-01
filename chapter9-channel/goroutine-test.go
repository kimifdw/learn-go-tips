package main

import (
	"runtime"
)

func main() {
	c := make(chan struct{})
	ci := make(chan int, 100)
	go func(i chan struct{}, j chan int) {
		// sum := 0
		// for i := 0; i < 10000; i++ {
		// 	sum += i
		// }
		// println(sum)
		// c <- struct{}{}
		// time.Sleep(1 * time.Second)
		for i := 0; i < 10; i++ {
			ci <- i
		}
		close(ci)

		c <- struct{}{}
	}(c, ci)

	println("NumGoroutine=", runtime.NumGoroutine())

	println("GOMAXPROCS=", runtime.GOMAXPROCS(0))

	//time.Sleep(5 * time.Second)

	<-c

	println("RunGoroutine=", runtime.NumGoroutine())
	for v := range ci {
		println(v)

	}
}
