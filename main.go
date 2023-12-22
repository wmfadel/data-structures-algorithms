package main

import (
	"data-structures-algorithms/stack"
	"fmt"
)

func main() {

	stack := stack.NewStack[string]()
	fmt.Println(stack.IsEmpty())

	stack.Push("A")
	stack.Push("B")
	stack.Push("C")
	stack.Push("D")
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.IsEmpty())

}
