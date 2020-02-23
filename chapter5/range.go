// 1. 遍历切片或映射
// 2. 第一个为下标，第二个为该下标所对应元素的副本
// 3. _: 忽略下标
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Println("2**%d=%d\n", i, v)
	}

	row := make([]int, 10)
	for i := range row {
		row[i] = 1 << unit(i) // == 2**i
	}

	for _, value := range row {
		fmt.Println("%d\n", value)
	}
}
