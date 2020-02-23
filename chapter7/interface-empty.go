package main

import "fmt"

type Inter interface {
	Ping()
	Pong()
}

type St struct{}

func (St) Ping() {
	println("Ping")
}

func (*St) Pong() {
	println("Pong")
}

func main() {
	var st *St = nil
	var it Inter = st

	fmt.Printf("%p\n", st)
	fmt.Printf("%p\n", it)

	if it != nil {
		it.Pong()
	}
}
