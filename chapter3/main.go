package main

import "fmt"

// 定义全局变量
var z string = "MAX"

func main() {
	x := "Hello,World" // 等同于var x string = "Hello,World"
	fmt.Println(x)
	//var y string = "Hello,World"
	var (
		y = "Hello World"
		z = "MAX2"
	) // 定义多个变量
	fmt.Println(x == y)

	//z := "MAX"
	fmt.Println("My cat's name is", z)

	const X string = "const Hello,World" // const:定义常量，不能被赋值
	fmt.Println(X)

	fmt.Println("Enter a number:")
	var input float64

	fmt.Scanf("%f", &input) // &input获取变量的值

	output := input * 2
	fmt.Println(output)

}