package main

import "fmt"

/**
定义平均数
**/
func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

/**
* 返回多个返回值，以逗号分隔
**/
func f() (int, int) {
	return 5, 6
}

/**
 * 绑定0个或多个参数
 **/
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

/**
* 函数中再使用函数
**/
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

// 传入参数为指针
func one(xPtr *int) {
	*xPtr = 1 // 将整数1存入xPtr对象指向的内存地址
}
func square(x *float64) float64 {
	*x = *x * *x
	return *x
}

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))

	x, y := f()
	fmt.Println("x,y", x, y)

	// 绑定多个参数
	fmt.Println(add(1, 2, 3))

	// 传入slice对象
	z := []int{1, 2, 3}
	fmt.Println(add(z...))

	// 将函数作为变量传入
	addNew := func(x, y int) int {
		return x + y
	}
	fmt.Println(addNew(1, 2))
	defer second()
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())

	// defer表示将second函数放到最后才执行
	// defer second()
	first()

	xPtr := new(int) // 声明整数对象
	one(xPtr)
	fmt.Println(*xPtr)

	t := 1.5
	fmt.Println(square(&t)) // &t返回*int
}
