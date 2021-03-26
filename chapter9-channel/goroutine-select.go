package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go func(chan int) {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			case <-time.After(1 * time.Second):
				fmt.Println("超时")
				break
			default:

			}
		}
	}(ch)

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second)
		println(<-ch)
	}
}
