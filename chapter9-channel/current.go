package main

import "fmt"

func main() {
	data := make([]int, 4)
	// channel参数只支持写入
	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i, _ := range data {
			handleData <- i
		}
	}

	// 传递channel
	handleData := make(chan int)
	go loopData(handleData)

	for datum := range handleData {
		fmt.Println(datum)
	}

	//***********************//

	resultData := func() <-chan int {
		results := make(chan int, 5)
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) {
		for result := range results {
			fmt.Printf("received:%d\n", result)
		}
		fmt.Println("done!")
	}

	results := resultData()
	consumer(results)

}
