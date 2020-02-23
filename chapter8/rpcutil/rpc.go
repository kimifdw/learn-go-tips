package rpcutil

import (
	"fmt"
	"net"
	"net/rpc"
)

type Server struct{}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

// 定义rpc的server
func ServerRpc() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":8899")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(c)
	}
}

// 定义rpc的客户端
func ClientRPC() {
	c, err := rpc.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int64
	err = c.Call("Server.Negate", int64(999), &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server.Negate(999) = ", result)
	}
}
