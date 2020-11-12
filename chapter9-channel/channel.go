package main

import "fmt"

import "time"

func addWithChannel(a, b int, ch chan int) {
	c := a + b
	fmt.Printf("%d + %d = %d\n", a, b, c)
	ch <- 1
}

// 只写通道
func test() <-chan int {
	ch := make(chan int, 20)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}

		close(ch)
	}()
	return ch

}

func main() {
	start := time.Now()
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go addWithChannel(1, i, chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
	end := time.Now()
	consume := end.Sub(start).Seconds()
	fmt.Println("program execute time:", consume)

	start2 := time.Now()
	ch2 := test()
	for i := range ch2 {
		fmt.Println("receive data:", i)
	}
	// for i := 0; i < len(ch2); i++ {

	// }
	end2 := time.Now()
	consume = end2.Sub(start2).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
