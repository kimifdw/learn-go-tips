package main

import (
	"fmt"
	"runtime"

	"time"
)

func sum(seq int, ch chan int) {
	defer close(ch)

	sum := 0
	for i := 0; i <= 10000000; i++ {
		sum += i
	}
	fmt.Printf("子协程%d运算结果：%d\n", seq, sum)
	ch <- sum
}

func main() {
	start := time.Now()
	cpuNum := runtime.NumCPU()
	runtime.GOMAXPROCS(1)
	chs := make([]chan int, cpuNum)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int, 1)
		go sum(i, chs[i])
	}
	sum := 0
	for _, ch := range chs {
		res := <-ch
		sum += res
	}
	end := time.Now()
	fmt.Printf("CPU NUM: %d,最终运算结果: %d, 执行耗时(s): %f\n", cpuNum, sum, end.Sub(start).Seconds())
}
