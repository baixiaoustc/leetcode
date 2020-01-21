package leetcode

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprint(n.Val)
}

func (n *ListNode) P() {
	if n == nil {
		return
	}
	fmt.Printf("node %s\n", n.String())
}

//n as the list head
func (n *ListNode) Print() {
	if n == nil {
		fmt.Println("nil list !!!")
		return
	}
	fmt.Println("print list")
	for p := n; p != nil; p = p.Next {
		fmt.Println(p.String())
	}
}

func (n *ListNode) Length() int {
	length := 0
	p := n
	for p != nil {
		p = p.Next
		length++
	}
	return length
}

//create list using []int
func createList(list []int) *ListNode {
	l := new(ListNode)
	p := l
	for i, e := range list {
		if i == 0 {
			p.Val = e
		} else {
			tmp := new(ListNode)
			tmp.Val = e
			p.Next = tmp
			p = p.Next
		}
	}

	return l
}

//find n th node of the list
//input: 1->2->3, n=0
//output: 1->2->3
//input: 1->2->3, n=1
//output: 2->3
func findList(head *ListNode, n int) *ListNode {
	if n < 1 {
		return head
	}

	p := head
	for ; n > 0; n-- {
		if p == nil {
			break
		}
		p = p.Next
	}

	return p
}

//cut previous n node of the list
//input: 1->2->3, n=0
//output: 1->2->3
//input: 1->2->3, n=1
//output: 2->3, notice: 1 -x-> 2
func cutList(head *ListNode, n int) *ListNode {
	if n < 1 {
		return head
	}

	for n -= 1; n > 0; n-- {
		if head == nil {
			break
		}
		head = head.Next
	}

	if head == nil {
		return nil
	}

	next := head.Next
	head.Next = nil
	return next
}

//reverse the list
//input: 1->2->3
//output: 3->2->1
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	tmp := new(ListNode)
	cur := head
	next := head.Next
	for next != nil {
		tmp = next.Next
		next.Next = cur
		cur = next
		next = tmp
	}

	head.Next = nil
	return cur
}
