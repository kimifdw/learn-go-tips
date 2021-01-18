// 1. 由方法签名定义的集合
// 2. 类型通过实现一个接口的所有方法来实现该接口
// 3. 隐式接口
// 4. 接口也是值,可作为函数的参数或返回值，保存了一个具体底层类型的具体值
// 5. 保存了nil具体值的接口其自身并不为nil
// 6. nil接口值，既不保存值也不保存具体类型
// 7. 空接口：指定了零个方法的接口值
// 8. 类型断言：访问接口值底层具体值的方法；返回其底层值和断言是否成功的布尔值
// 9. 类型选择：一种按顺序从几个类型断言中选择分支的结构
// 10. Stringer是一个可以用字符串描述自己的类型
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type I interface {
	M()
}

type Reader interface {
	Read(p []byte) (n int, err error)
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f
	a = &v

	a = v

	fmt.Println(a.Abs())

	var i T = T{"hello"}
	i.M()

	// 空接口值
	var k interface{}

	var i interface{} = "hello"
	// 类型断言
	s := i.(string)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
