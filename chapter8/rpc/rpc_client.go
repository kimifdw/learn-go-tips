package main

import (
	// "chapter8/rpc/utils"
	"log"
    // "net/rpc"
    // "fmt"
    "net"
    "net/rpc/jsonrpc"
    "time"
)

func main() {
	// var serverAddress = "localhost"
	// client, err := rpc.DialHTTP("tcp", serverAddress+":8080")
	// if err != nil {
	// 	log.Fatal("建立服务失败：", err)
	// }

    // args := &utils.Args{10, 10}
    // var reply int
    // err = client.Call("MathService.Multiply", args, &reply)
    // if err != nil {
    //     log.Fatal("Multiply调用失败：", err)
    // }
    // fmt.Printf("%d*%d=%d\n",args.A, args.B,reply)

    conn, err := net.DialTimeout("tcp", "localhost:8081", 30*time.Second)
    if err != nil {
        log.Fatalf("client connect failed: %v",err)
    }
    defer conn.Close()

    client := jsonrpc.NewClient(conn)
    var item Item
    client.Call("ServiceHandler.GetName", 1, &item)
    log.Printf("ServiceHandler.GetName return result: %v\n", item)

    var resp Response
    item = Item{2, "Emily"}
    client.Call("ServiceHandler.SaveName", item, &resp)
    log.Printf("ServiceHandler.SaveName return result: %v\n", resp)
}
