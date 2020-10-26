package main

import (
	"fmt"
	"strings"
)

/**
给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。
例如：haystack="hello",needle="ll",结果为2
*/
func strStr(haystack string, needle string) int {

	needleLen := len(needle)
	haystackLen := len(haystack)

	// 处理needle字符串为空的情况
	if needleLen == 0 {
		return 0
	}

	// 处理其中一个字符串为空气的情况
	if haystackLen == 0 || haystackLen < needleLen {
		return -1
	}

	var i int
	for i = 0; i < haystackLen-needleLen+1; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
		//for j = 0; j < needleLen; j++ {
		//	// 判断数组是否越界的问题
		//	z := i + j
		//	if z > haystackLen-1 {
		//		z = haystackLen - 1
		//	}
		//	if haystack[z] != needle[j] {
		//		break
		//	}
		//}
		//
		//if len(needle) == j {
		//	return i
		//}
	}
	return -1

}

func strStr1(havstack string, needle string) int {
	return strings.Index(havstack, needle)
}

func main() {
	//haystack := "mississippi"
	//needle := "issipi"
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStr1(haystack, needle))
}
