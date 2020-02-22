package leetcode

import (
	"fmt"
	"testing"
)

func TestLinkedQueue1(t *testing.T) {
	queue := NewLinkedQueue()
	queue.Add(1)
	fmt.Println(queue.Peek())
	queue.Remove()
	queue.Print()
	queue.Add(2)
	queue.Add(3)
	fmt.Println(queue.Peek())
	queue.Remove()
	fmt.Println(queue.Peek())
	fmt.Println(queue.Size())
}

func TestLinkedQueue2(t *testing.T) {
	queue := NewLinkedQueue()
	queue.Add([2]int{1, 2})
	fmt.Println(queue.Peek())
	fmt.Println(queue.Size())
}
