package main

import "fmt"

func main() {
	fmt.Println("1+1=", 1.0 + 1.0)
	fmt.Println(len("Hello,World")) // len读取字符串长度
	fmt.Println("Hello,World"[1]) // 字符串数组的index从0开始，返回的是字节
	fmt.Println("Hello," + "world") // +默认将两个字符串串联起来

}
