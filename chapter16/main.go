package main

import "fmt"

import "time"

const LIM = 41

var fibs [LIM]uint64

func main() {
	// result := 0
	var resultNew uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		resultNew = fibonacciNew(i)
		fmt.Printf("fibonacciNew(%d) is: %d\n", i, resultNew)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("long Calculation took this amount of time: %s\n", delta)
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return res
}

func fibonacciNew(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return res
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacciNew(n-1) + fibonacciNew(n-2)
	}
	fibs[n] = res
	return res
}
