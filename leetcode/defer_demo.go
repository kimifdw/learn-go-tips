package main

func main() {
	defer func() {
		println(`defer 1`)
	}()
	defer func() {
		println(`defer 2`)
	}()
}
