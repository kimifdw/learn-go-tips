package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

type Op func(int, int) int // 定义函数类型

func do(f Op, a, b int) int {
	return f(a, b)
}

func main() {
	a := do(add, 1, 2)
	fmt.Println(a)

	s := do(sub, 2, 1)
	fmt.Println(s)
}
