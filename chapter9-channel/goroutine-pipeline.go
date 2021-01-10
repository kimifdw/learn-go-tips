package main

import "fmt"

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func gen(a, b int) <-chan int {
	in := make(chan int)
	go func() {
		in <- a
		in <- b
		close(in)
	}()

	//defer close(in)
	return in
}

func main() {
	c := gen(2, 3)
	out := sq(c)
	fmt.Println(<-out)
	fmt.Println(<-out)
}
