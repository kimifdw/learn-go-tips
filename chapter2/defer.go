// 1. 推迟函数到外层函数返回之后执行
// 2. 函数会立即求值，但只在外层函数返回后才执行，按后进先出的顺序调用
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
