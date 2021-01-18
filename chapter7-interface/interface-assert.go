package main

import "fmt"

type Inter interface {
	Ping()
	Pang()
}

type Anter interface {
	Inter
	String()
}

type St struct {
	Name string
}

func (St) Ping() {
	println("ping")
}

func (*St) Pang() {
	println("Pang")
}

func main() {
	st := &St{"andes"}

	var i interface{} = st

	if o, ok := i.(Inter); ok {
		o.Ping()
		o.Pang()
	}

	if p, ok := i.(Anter); ok {
		p.String()
	}

	if s, ok := i.(*St); ok {
		fmt.Printf("%s", s.Name)

	}
}
