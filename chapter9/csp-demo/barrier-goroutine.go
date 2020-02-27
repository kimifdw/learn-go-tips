// Barrier模式 阻塞直到聚合所有goroutine返回结果
// 场景 多个网络请求并发，聚合结果；粗粒度任务拆分并发执行，聚合结果
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

//BarrierResp 响应结果
type BarrierResp struct {
	Err    error
	Resp   string
	Status int
}

func makeRequest(out chan<- BarrierResp, url string) {
	res := BarrierResp{}

	client := http.Client{
		Timeout: time.Duration(2 * time.Second),
	}

	resp, err := client.Get(url)
	if resp != nil {
		res.Status = resp.StatusCode
	}
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	byte, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		res.Err = err
		out <- res
		return
	}

	res.Resp = string(byte)
	out <- res
}

// barrier 结果合并
func barrier(endpoints ...string) {
	requestNum := len(endpoints)

	in := make(chan BarrierResp, requestNum)
	response := make([]BarrierResp, requestNum)

	defer close(in)

	for _, endpoint := range endpoints {
		go makeRequest(in, endpoint)
	}

	var hasError bool
	for i := 0; i < requestNum; i++ {
		resp := <-in
		if resp.Err != nil {
			fmt.Println("ERROR:", resp.Err, resp.Status)
			hasError = true
		}
		response[i] = resp
	}
	if !hasError {
		for _, resp := range response {
			fmt.Println(resp.Resp)
		}
	}
}

// 利用锁实现
func barrierWithLock(endpoints ...string) {
	var g errgroup.Group
	var mu sync.Mutex

	response := make([]BarrierResp, len(endpoints))

	for i, endpoint := range endpoints {
		i, endpoint := i, endpoint
		g.Go(func() error {
			res := BarrierResp{}
			resp, err := http.Get(endpoint)
			if err != nil {
				return err
			}

			byt, err := ioutil.ReadAll(resp.Body)
			defer resp.Body.Close()

			if err != nil {
				return err
			}

			res.Resp = string(byt)
			res.Status = resp.StatusCode
			mu.Lock()
			response[i] = res
			mu.Unlock()
			return err
		})

		if err := g.Wait(); err != nil {
			fmt.Println(err)
		}

		for _, resp := range response {
			fmt.Println(resp.Status)
		}

	}
}

func main() {
	barrierWithLock([]string{"http://www.baidu.com", "http://www.sina.com", "https://cn.bing.com"}...)
}
