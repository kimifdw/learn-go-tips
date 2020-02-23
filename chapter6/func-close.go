package main

import "fmt"

func fa(a int) func(i int) int {
	return func(i int) int {
		println(&a, a, i)
		a = a + i
		return a
	}
}

func main() {
	f := fa(1)
	g := fa(1)

	fmt.Println(f(1))
	fmt.Println(f(1))

	fmt.Println(g(1))
}
