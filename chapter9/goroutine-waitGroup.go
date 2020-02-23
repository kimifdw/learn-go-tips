package main

import (
	"net/http"
	"sync"
)

var wg sync.WaitGroup
var urls = []string{
	"http://www.sina.com.cn",
	"http://www.golang.org/",
	"http://www.baidu.com/",
}

func main() {
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
}
