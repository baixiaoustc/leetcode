package leetcode

import (
	"fmt"
	"testing"
)

func TestStackInterface(t *testing.T) {
	stack := NewStackInterface()
	stack.Push("abc")
	v1 := stack.Peak()
	fmt.Println(v1)
	stack.Push(66)
	v2 := stack.Pop()
	fmt.Println(v2)
	switch v2.(type) {
	case int:
		fmt.Println("int")
	}
}
