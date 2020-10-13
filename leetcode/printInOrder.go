// https://leetcode-cn.com/problems/print-in-order/?utm_source=LCUS&utm_medium=ip_redirect_q_uns&utm_campaign=transfer2china
// https://zhuanlan.zhihu.com/p/103405131
package main

import (
	"fmt"
	"strconv"
)

func First(streamSync [3]chan interface{}) {
	fmt.Print("First ")
	streamSync[0] <- nil
}

func Second(streamSync [3]chan interface{}) {
	<-streamSync[0]
	fmt.Print("Second ")
	streamSync[1] <- nil
}

func Third(streamSync [3]chan interface{}) {
	<-streamSync[1]
	fmt.Print("Third")
	streamSync[2] <- nil
}

// 顺序打印
func printInOrder(callOrder [3]int) {
	inputCallOrder := callOrder
	fmt.Println("[] inputCallOrder:", inputCallOrder)

	var streamSync [3]chan interface{}
	for i := range streamSync {
		streamSync[i] = make(chan interface{})
	}

	var functionNumTable = map[int]func([3]chan interface{}){
		1: First,
		2: Second,
		3: Third,
	}

	for _, fNum := range inputCallOrder {
		go functionNumTable[fNum](streamSync)
	}

	<-streamSync[2]
}

//func main() {
//   var testCases = [][3]int{
//       {1, 2, 3},
//       {1, 3, 2},
//       {2, 1, 3},
//       {2, 3, 1},
//       {3, 1, 2},
//       {3, 2, 1},
//   }
//
//   for _, theCase := range testCases {
//       printInOrder(theCase)
//       fmt.Println()
//       fmt.Println()
//   }
//}

func main() {
	var m int32 = 0x12345678
	var n int8 = int8(m)
	str := strconv.FormatInt(int64(n), 16)
	bytes := Hextob(str)
	h := ByteToH(bytes)
	fmt.Println(h)
}

func ByteToH(b []byte) (H string) {
	H = fmt.Sprintf("%x", b)
	return
}

func Hextob(str string) []byte {
	slen := len(str)
	bHex := make([]byte, len(str)/2)
	ii := 0
	for i := 0; i < len(str); i = i + 2 {
		if slen != 1 {
			ss := string(str[i]) + string(str[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			bHex[ii] = byte(bt)
			ii = ii + 1
			slen = slen - 2
		}
	}
	return bHex
}
