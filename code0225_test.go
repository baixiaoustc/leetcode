package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructorMyStack(t *testing.T) {
	stack := ConstructorMyStack()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	stack.Push(3)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
