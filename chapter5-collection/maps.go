// 1. 将键映射到值
// 2. 零值为nil.
// 3. 忽略顶级类型
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	// m = make(map[string]Vertex)
	// m["Bell Labs"] = Vertex{
	// 	40.68433, -74.39967,
	// }
	fmt.Println(m)

	n := make(map[string]int)
	n["Answer"] = 42
	fmt.Println("The value:", n["Answer"])

	n["Answer"] = 48
	fmt.Println("The value:", n["Answer"])

	// 删除
	delete(n, "Answer")
	fmt.Println("The value:", n["Answer"])

	// 通过双赋值检测某个键是否存在
	v, ok := n["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
