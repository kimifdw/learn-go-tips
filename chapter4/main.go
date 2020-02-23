package main

import "fmt"

func main() {
	fmt.Println(`1 2 3 4 5 6 7 8 9 10`) // 等同于fmt.Println(1),fmt.Println(2)将数据一个个打印出来
	i := 1
	for i <= 10 {
		fmt.Println("i=", i)
		i += 1
	} // for循环

	// for循环+if条件
	for i := 11; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Println("even=", i)
		} else if i % 3 == 0 {
			fmt.Println("even3=", i)
		} else {
			fmt.Println("odd=", i)
		}
	}

	// for循环+switch条件
	for i := 21; i <= 30; i++ {
		switch i {
		case 21: fmt.Println("21")
		case 22: fmt.Println("22")
		default: fmt.Println("Unknown Number")
		}
	}
}
