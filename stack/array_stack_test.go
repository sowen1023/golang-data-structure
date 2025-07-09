package stack_test

import (
	"fmt"
	"testing"

	"io.github.sowen.datastructure/stack"
)

func TestArrayStack_Push(t *testing.T) {
	s := stack.NewWithCapacity[int](8)

	for i := 1; i <= 10; i++ {
		s.Push(i)
		fmt.Println(s)
	}

	fmt.Printf("Top element = %v\n", s.Peek())
	fmt.Printf("Is empty = %v\n", s.IsEmpty())

	for i := 1; i <= 10; i++ {
		s.Pop()
		fmt.Println(s)
	}

	fmt.Printf("Is empty = %v\n", s.IsEmpty())
}
