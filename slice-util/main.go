package main

import "fmt"

func a() []int {
	a1 := []int{3}
	a2 := a1[1:]
	return a2
}

func b() []int {
	a1 := []int{3, 4, 5, 6, 7, 8}
	a2 := a1[2:]
	return a2
}

func main() {
	fmt.Println(b())
}
