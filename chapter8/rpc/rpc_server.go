package main

import (
	"errors"
	"log"
	"net"
	// "net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// MathService 计算服务
type MathService struct{}

// ServiceHandler 服务
type ServiceHandler struct{}

// Multiply 乘法
func (m *MathService) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Divide 除法
func (m *MathService) Divide(args *Args, reply *int) error {
	if args.B == 0 {
		return errors.New("除数不能为0")
	}
	*reply = args.A / args.B
	return nil
}

// GetName 获取名称
func (serviceHandler *ServiceHandler) GetName(id int, item *Item) error {
	log.Printf("receive GetName call, id: %d", id)
	item.ID = id
	item.Name = "Emily"
	return nil
}

// SaveName 保存名称
func (serviceHandler *ServiceHandler) SaveName(item *Item, resp *Response) error {
	log.Printf("receive SaveName call, Item: %v", item)
	resp.Ok = true
	resp.ID = item.ID
	resp.Message = "success"
	return nil
}

func main() {

	// math := new(MathService)

	// // 注册服务
	// rpc.Register(math)
	// // 以HTTP服务作为rpc的服务器
	// rpc.HandleHTTP()
	// // 指定8080端口
	// listener, rpcErr := net.Listen("tcp", "localhost"+":8080")
	// if rpcErr != nil {
	// 	log.Fatal("启动服务失败：", rpcErr)
	// }
	// // 启动服务
	// rpcErr = http.Serve(listener, nil)
	// if rpcErr != nil {
	// 	log.Fatal("启动HTTP服务失败：", rpcErr)
	// }

	server := rpc.NewServer()
	jsonListener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("listening port error: %v", err)
	}
	defer jsonListener.Close()

	log.Println("start listen on port localhost:8081")

	serviceHandler := &ServiceHandler{}
	err = server.Register(serviceHandler)
	if err != nil {
		log.Fatalf("register failed: %v", err)
	}

	for {
		conn, err := jsonListener.Accept()
		if err != nil {
			log.Fatalf("receive failed:%v", err)
		}
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
