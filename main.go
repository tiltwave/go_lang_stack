package main

import (
	"./stack"
	"fmt"
)

type MyDataType struct {
	i int
	s string
}

func main() {
	s := stack.Stack{}
	s.Push(1)
	s.Push(10.5)
	s.Push("hello")
	s.Push([]string{"hello", "bye"})
	s.Push([]int{200, 300})
	s.Push(MyDataType{1, "hey"})

	fmt.Println("Number of items stored:", s.Len())

	for !s.IsEmpty() {
		item, _ := s.Pop()
		fmt.Printf("Data type: %T\n", item)
		fmt.Printf("Data value: %v\n\n", item)
	}
}
