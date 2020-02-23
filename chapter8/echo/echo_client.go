package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
)

var host = flag.String("host", "localhost", "host")
var port = flag.String("port", "3333", "port")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *host+":"+*port)
	if err != nil {
		fmt.Println("Error connnecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connection to " + *host + ":" + *port)

	var wg sync.WaitGroup

	wg.Add(2)
	go handleWrite(conn, &wg)
	go handleRead(conn, &wg)

	wg.Wait()
}

func handleWrite(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 10; i > 0; i-- {
		_, e := conn.Write([]byte("hello " + strconv.Itoa(i) + "\r\n"))
		if e != nil {
			fmt.Println("Error to send message because of ", e.Error())
			break
		}
	}

}

func handleRead(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	buf := make([]byte, 1024)
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error to read message because of ", err)
		return
	}
	fmt.Println(string(buf[:reqLen-1]))
}
