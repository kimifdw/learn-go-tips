package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go http.ListenAndServe("127.0.0.1:6060", nil)
	for {
		b := make([]byte, 4096)
		for i := 0; i < len(b); i++ {
			b[i] = b[i] + 0xf
		}
		time.Sleep(time.Nanosecond)
	}
}
