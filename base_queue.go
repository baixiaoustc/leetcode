package leetcode

import "fmt"

//////////////////////////////////// FIFO QUEUE //////////////////////////////////////

type QueueNode struct {
	Val  interface{}
	Next *QueueNode
	Pre  *QueueNode
}

//FIFO
type LinkedQueue struct {
	head *QueueNode
	tail *QueueNode
	size int
}

func (q *LinkedQueue) Add(val interface{}) {
	node := &QueueNode{Val: val}
	if q.size == 0 {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		q.tail = node
	}
	q.size++
}

func (q *LinkedQueue) Peek() (interface{}, error) {
	if q.head == nil || q.size == 0 {
		return -1, fmt.Errorf("nil queue")
	}
	return q.head.Val, nil
}

func (q *LinkedQueue) Remove() error {
	if q.head == nil || q.size == 0 {
		return fmt.Errorf("nil queue")
	}
	first := q.head
	q.head = first.Next
	first = nil
	q.size--
	return nil
}

func (q *LinkedQueue) Size() int {
	return q.size
}

func (q *LinkedQueue) Print() {
	tmp := q.head
	fmt.Println("print queue")
	for tmp != nil {
		fmt.Println(tmp.Val)
		tmp = tmp.Next
	}
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}
