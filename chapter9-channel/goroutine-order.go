package main

import "fmt"

func main() {
	// 抛异常
	//a := "hello world"
	//go func() {
	//    fmt.Println(a)
	//    a = "hello goroutine"
	//    go func() { fmt.Println(a) }()
	//}()
	//
	//select {}

	var a = "hello"
	go func() {
		fmt.Println(a)
	}()
	go func() {
		fmt.Println(a)
	}()

	a = "world"

	select {}
}
