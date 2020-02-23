package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T values: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T values: %v\n", MaxInt, ToBe)
	fmt.Printf("Type: %T values: %v\n", ToBe, ToBe)
}
