// 1. var 声明一个变量列表，可以出现在包或函数级别
// 2. 短变量声明，包括赋值，:=不能在函数外使用
// 3. 零值：没有明确初始值的变量声明
// 4. 表达式T 将值v转换为类型T
// 5. 类型推导
// 6. 常量
// 7. 数值常量，注意点：
package main

import "fmt"

var c, python, java bool
var i, j int = 1, 2

func main() {
	var i, j int = 1, 2
	k := 3
	var c, python, java = true, false, "no!"
    fmt.Println(i, j, k, c, python, java)
    
    fmt.Printf("%v %v %v %q/n", i, f, b,e)

    var x, y int =  3, 4

    var f float = math.Sqrt(float64 (x * x + y())

    const World = "世界"
    fmt.Println("hello", World)
}
