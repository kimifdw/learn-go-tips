// 1. 没有类。允许为结构体类型定义方法
// 2. 方法是一类带特殊的接收者参数的函数
// 3. 方法即函数，只是带接收者参数的函数
// 4. 只能为在同一包内定义的类型的接收者声明方法，不能为其它包内定义的类型的接收者声明方法
// 5. 指针接收者，可以修改接收者指向的值
// 6. 以指针为接收者的方法被调用时，接收者既能为值又能为指针
// 7. 避免每次调用方法时复制该值
package main

import (
    "fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	// fmt.Println(v.Abs())
	v.Scale(2)
	ScaleFunc(&v, 10)

	fmt.Println(Abs(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	Scale(&v, 10)
	fmt.Println(Abs(v))

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
