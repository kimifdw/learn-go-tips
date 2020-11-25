// 1. 结构体：一组字段
// 2. 使用点号访问
// 3. 通过结构体指针访问
// 4. 直接列出字段的值来新分配一个结构体
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {

	// v := Vertex{1, 2}
	// v.X = 4
	// p := &v
	// p.X = 1e9
	// fmt.Println(v)

	fmt.Println(v1, p, v2, v3)

}
