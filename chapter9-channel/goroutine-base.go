package main

import (
	"fmt"
	"math/rand"
)

func GenerateIntABase(done chan struct{}) chan int {
	ch := make(chan int, 10)

	go func() {
	Lable:
		for {
			select {
			case ch <- rand.Int():
			case <-done:
				break Lable
			}

		}
		close(ch)
	}()
	return ch
}

func GenerateIntB() chan int {
	ch := make(chan int, 10)
	go func() {
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func GenerateInt() chan int {
	done := make(chan struct{})
	ch := make(chan int, 20)
	go func() {
		select {
		case ch <- <-GenerateIntABase(done):
		case ch <- <-GenerateIntB():
		}
	}()
	return ch
}

func main() {
	done := make(chan struct{})
	ch := GenerateIntABase(done)
	for i := 0; i < 100; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	close(ch)
}
