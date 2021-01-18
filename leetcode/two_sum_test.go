package main

import (
	"fmt"
	"testing"
)

func Test_twoSum(t *testing.T) {
	sum := twoSum([]int{2, 11, 7, 15}, 9)

	fmt.Printf("result:%d\n", sum)
}
