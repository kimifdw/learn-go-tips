package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"sync"
	"time"
)

// doSomething：模拟抓取任务的执行
func doSomething(u string) {
	fmt.Println(u)
	time.Sleep(2 * time.Second)
}

const (
	// 同时并行运行的goroutine上限
	Limit = 3
	// 每个goroutine获取信号量资源的权重
	Weight = 1
)

func main() {
	urls := []string{
		"http://www.baidu.com",
		"http://www.sina.com",
		"http://www.google.com",
	}
	s := semaphore.NewWeighted(Limit)

	var w sync.WaitGroup
	for _, url := range urls {
		w.Add(1)
		go func(u string) {
			s.Acquire(context.Background(), Weight)
			doSomething(u)
			s.Release(Weight)
			w.Done()
		}(url)
	}
	w.Wait()
	fmt.Println("All Done")
}
