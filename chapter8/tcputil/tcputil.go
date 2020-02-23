package tcputil

import (
	"encoding/gob"
	"fmt"
	"net"
)

func Server() {
	// 监听端口
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		// 等待下次调用并返回一个Connection对象
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleServerConnection(c)
	}
}

/**
* net.Conn：面向流的网络连接对象
**/
func handleServerConnection(c net.Conn) {
	var msg string
	// 将消息进行解码并将解码的内容存放到msg中
	err := gob.NewDecoder(c).Decode(&msg)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("接收到的信息", msg)
	}
	c.Close()
}

// TCP客户端
func Client() {
	// Dial：以tcp方式连接端口为9999的server
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	msg := "Hello world!"
	fmt.Println("发送的消息：", msg)
	// 将消息编码并发送给server
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}
