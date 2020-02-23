package main

import (
	"fmt"
	"sync"
)

func testPut(pool *sync.Pool, deferFunc func()) {
	defer func() {
		deferFunc()
	}()
	value := "Hello, Emily"
	pool.Put(value)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello world!"
		},
	}
	go testPut(pool, wg.Done)
	// value := "Hello, Emily"
	// pool.Put(value)
	wg.Wait()
	fmt.Println(pool.Get())
}
