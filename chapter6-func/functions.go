// 1. 类型在变量名之后，连续多个函数的已命名形参类型相同，可进行类似addNew方法的参数设定
// 2. 支持返回任意数量的值，参考swap函数
// 3. 支持返回值命名，参考split函数，没有参数的return语句返回已命名的返回值
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func addNew(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(addNew(11, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))
}
