// 1. 保存值的内存地址
// 2. 类型*T是指向T类型值的指针，其零值为nil
// 3. & [会生成一个指向其操作数的指针]
// 4. * [指针指向的底层的值]
package main

import "fmt"

func main() {
	i, j := 42, 2701
	// 指向i
	p := &i
	fmt.Println(*p)
	// 通过指针设置i的值
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

}
