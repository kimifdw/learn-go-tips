package main

import (
    "fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/**
 * 实现goroutines的demo
 **/
func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

/**
 * 实现channel的demo
 * 发送器
 **/
func pinger(c chan string) {
	for i := 0; ; i++ {
		// <-：从channel中发送和接收消息。
		// 发送ping到channel
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

/**
 * 接收器
 **/
func printer(c chan string) {
	for {
		// 从channel中接收消息并将值复制给msg。
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	go f(0) // 创建一个并发routine

	c := make(chan string)

	go pinger(c)

	go ponger(c)

	go printer(c)
	// 初始化channel
	c1 := make(chan string)
	c2 := make(chan string)
	// 每2秒输出from 1
	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()
	// 每3秒输出from 2
	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			// 随机从任意一个channel钟接收到消息并打印出来
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				// default:
				// 	fmt.Println("nothing is ready")
			}
		}
	}()

	urls := []string{
		"http://www.apple.com",
		"http://www.sina.com.cn",
		"http://www.baidu.com",
		"http://www.microsoft.com",
	}
	// 初始化基于homHomePageSize对象的channel
	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			// 读取url
			res, err := http.Get(url)
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()

			bs, err := ioutil.ReadAll(res.Body)
			if err != nil {
				panic(err)
			}

			// 将HomHomePageSize对象丢入channel
			fmt.Println("send ", url, ":", len(bs))
			results <- HomePageSize{
				URL:  url,
				Size: len(bs),
			}
		}(url)
	}

	var biggest HomePageSize

	for _, url := range urls {
		// 读取channel中的内容(无序的)
		time.Sleep(time.Second * 2)
		result := <-results
		fmt.Println("receive", url, ":", result)
		if result.Size > biggest.Size {
			biggest = result
		}

	}

	fmt.Println("The biggest home page:", biggest.URL)

	var input string

	fmt.Scanln(&input)

}
