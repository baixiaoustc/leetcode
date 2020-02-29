package leetcode

import (
	"fmt"
	"testing"
)

func TestConstructorMyQueue(t *testing.T) {
	queue := ConstructorMyQueue()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
