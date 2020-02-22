package leetcode

import (
	"container/ring"
	"fmt"
	"testing"
)

func TestMyCircularQueue1(t *testing.T) {
	queue := Constructor(3)
	fmt.Println(queue.IsFull(), queue.IsEmpty())
	fmt.Println(queue.EnQueue(1), queue.IsEmpty())
	fmt.Println(queue.Front(), queue.Rear())
	fmt.Println(queue.DeQueue(), queue.IsEmpty())
	fmt.Println(queue.EnQueue(1), queue.EnQueue(2), queue.EnQueue(3))
	fmt.Println(queue.Front(), queue.Rear(), queue.IsFull())
	fmt.Println(queue.DeQueue(), queue.IsFull())
	fmt.Println(queue.Front(), queue.Rear())
}

func TestMyCircularQueue2(t *testing.T) {
	queue := ring.New(3)
	fmt.Println(queue.Len())
	queue.Value = 1
	fmt.Println(queue.Value)
	fmt.Println(queue.Next())
	queue.Value = 2
	fmt.Println(queue.Value)
	fmt.Println(queue.Next())
	queue.Value = 3
	fmt.Println(queue.Value)
	fmt.Println(queue.Next())
}
