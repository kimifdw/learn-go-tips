package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		fmt.Printf("k:%d;v:%d;\n", k, v)
		if idx, ok := m[target-v]; ok {
			return []int{idx, k}
		}
		m[v] = k
	}
	return nil
}

func main() {
	sum := twoSum([]int{2, 11, 7, 15}, 9)

	fmt.Printf("result:%d\n", sum)
}
