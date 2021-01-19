package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var (
	Cnt int
	mu  sync.Mutex
)

func Add(iter int) {
	mu.Lock()
	for i := 0; i < iter; i++ {
		Cnt++
	}
	mu.Unlock()
}

// go run main.go 2>trace.out ----输出out文件
// go tool trace trace.out --- 浏览器查看trace viewer
func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			Add(100000)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(Cnt)
}
