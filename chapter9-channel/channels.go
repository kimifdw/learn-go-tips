// 1. 带有类型的管道
// 2. <-: 将值发送到信道
// 3. <-ch：接收从信道返回的值
// 4. make(chan int)：信道必须使用前创建
// 5. 默认情况下，发送和接收操作在另一端准备好之前都会是阻塞
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
