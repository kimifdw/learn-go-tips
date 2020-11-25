package main

type Printer interface {
	Print()
}

type S struct{}

func (s S) Print() {
	println("print")
}

func main() {
	var i Printer

	i = S{}

	i.Print()
}
