package main

import "fmt"

type Stack struct {
	data []interface{}
}

// interface{} 类似于java的Object，可以被所有的Go类型实现
func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

// s *Stack：声明一个方法调用者s，类似于java的this
func (s *Stack) Pop() interface{} {
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil
	s.data = s.data[:i]
	return res
}

func (s *Stack) Size() int {
	return len(s.data)
}

func main() {
	var s Stack
	s.Push("World")
	s.Push("hello,")
	for s.Size() > 0 {
		fmt.Println(s.Pop())
	}
	fmt.Println()
}
