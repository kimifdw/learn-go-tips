package main

import "fmt"

func chain(in chan int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- 1 + v
		}
		close(out)
	}()
	return out
}

// 合理的写法，尽量缩小channel的能力
func chainNew() {
	// 只读channel
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5)
		go func() {
			defer close(resultStream)
			for i := 0; i < 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func main() {
	//in := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 10; i++ {
	//		in <- i
	//	}
	//	close(in)
	//}()
	//
	//out := chain(chain(in))
	//for v := range out {
	//	fmt.Println(v)
	//}
	chainNew()
}
