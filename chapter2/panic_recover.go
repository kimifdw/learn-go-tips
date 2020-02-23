package main

import "fmt"

// slice越界捕获
func sliceFoo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var bar = []int{1}
	fmt.Println(bar[1])
}

// 未初始化指针就在使用
func pointerFoo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var bar *int
	fmt.Println(*bar)
}

// 往关闭的channel写数据
func channelFoo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var bar = make(chan int, 1)
	close(bar)
	bar <- 1
}

// map并发读写
func mapFoo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var bar = make(map[int]int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		for {
			_ = bar[1]
		}
	}()

	for {
		bar[1] = 1
	}
}

// interface类型断言
func interfaceFoo() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var i interface{} = "abc"
	// 类型断言
	_ = i.([]string)
}

func main() {
	sliceFoo()
	pointerFoo()
	channelFoo()
	// mapFoo()
	interfaceFoo()
	fmt.Println("exit")
}
