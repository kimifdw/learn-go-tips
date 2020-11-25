package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4

	// 反射1
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(&x)

	fmt.Println("type:", t)

	fmt.Println("value:", v.String())
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	//fmt.Println("value:", v.Float())

	// 反射法则2
	o := v.Interface().(float64)
	fmt.Println("从反射得到：", o)

	// 反射法则3
	p := v.Elem()
	fmt.Println(p.CanSet())
	p.SetFloat(7.1)
	fmt.Println("法则3：", v)
}
