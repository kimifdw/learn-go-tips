package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.sina.com.cn",
	"http://www.golang.org/",
	"http://www.baidu.com/",
}

func main() {
	// 协程开启
	for _, url := range urls {

		// 每个url启动一个goroutine，同时给wg加1
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err == nil {
				println(resp.Status)
			}
		}(url)
	}
	wg.Wait()

	// 协程关闭
	wg.Add(1)
	go func() {
		var skip int
		for {
			_, file, line, ok := runtime.Caller(skip)
			if !ok {
				break
			}
			// 输出堆栈信息
			fmt.Printf("%s:%d\n", file, line)
			skip++
		}
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		// 退出协程
		runtime.Goexit()

		fmt.Println("never executed")
	}()
	wg.Wait()
}
