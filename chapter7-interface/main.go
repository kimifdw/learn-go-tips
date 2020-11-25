package main

import (
	"fmt"
	"math"
)

/**
 * 圆的面积
 **/
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

/**
 * 定义接口
 **/
type Shape interface {
	area() float64
}

func totalArea(shaps ...Shape) float64 {
	var area float64
	for _, value := range shaps {
		area += value.area()
	}
	return area
}

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, shape := range m.shapes {
		area += shape.area()
	}
	return area
}

/**
* 定义矩形类
**/
type Rectangle struct {
	x1, x2, y1, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

/**
* 矩形的面积
**/
func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

/**
 * 定义Circle类
 **/
type Circle struct {
	x, y, r float64
}

/**
* area方法能够被struct调用
**/
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Person struct {
	Name string
}

/**
 * struct包含子struct
 **/
type Android struct {
	Person
	Model string
}

func (p Person) Talk() {
	fmt.Println("hi,my name is ", p.Name)
}

func main() {

	c := Circle{x: 0, y: 0, r: 5} // 初始化Circle,指针的默认值为nil;&Circle表示为Circle对象的指针
	r := Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}

	fmt.Println(r.area())
	// fmt.Println(circleArea(&c))
	fmt.Println(c.area()) //

	android := new(Android)

	person := Person{Name: "测试"}
	person.Talk()
	android.Name = "测试2"
	android.Talk()

	fmt.Println(totalArea(&c, &r)) // 需要考虑同样的参数和方法名

	multishape := MultiShape{
		shapes: []Shape{
			Circle{0, 0, 5},
			Rectangle{x1: 0, y1: 0, x2: 10, y2: 10},
		},
	}
	fmt.Println(multishape.area())
}
