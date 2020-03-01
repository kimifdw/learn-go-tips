package main

import (
	"fmt"
	"sync"
)

// gen 生成
func gen(nums ...int) <-chan int {
	out := make(chan int, len(nums))
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// sq 将数据输入乘2
func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// merge 合并chan
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 1)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}

	}
	wg.Add(len(cs))

	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)
	c := gen(2, 3)
	c1 := sq(done, c)
	c2 := sq(done, c)

	out := merge(done, c1, c2)
	fmt.Println(<-out)

}