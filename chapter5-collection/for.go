// 1. 只有一种循环结构
// 2. 三个构成部分外没有小括号
package main

import "fmt"

func main() {
	sum := 1
	// 1. 标准for循环
	// for i := 0; i < 10; i++ {
	// 	sum += i
	// }

	// 2. 初始化和后置语句可选，类似while
	// for sum < 1000 {
	// 	sum += sum
	// }

	// 3. 省略条件即为无限循环
	// for {

	// }

	fmt.Println(sum)
}
