package leetcode

import (
	"testing"
)

type handle func(head *ListNode) *ListNode

var f handle = insertionSortListRecursive

func TestInsertionSortList1(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := f(l1)
	l2.Print()
}

func TestInsertionSortList3(t *testing.T) {
	list1 := []int{1, 3, 3, 7, 9, 11, 9}
	l1 := createList(list1)

	l2 := f(l1)
	l2.Print()
}

func TestInsertionSortList4(t *testing.T) {
	list1 := []int{1, 3, -5, 7, 0, 11}
	l1 := createList(list1)

	l2 := f(l1)
	l2.Print()
}

func TestInsertionSortList5(t *testing.T) {
	list1 := []int{4, 2, 1, 3}
	l1 := createList(list1)

	l2 := f(l1)
	l2.Print()
}

func TestInsertionSortList6(t *testing.T) {
	list1 := []int{4, 2}
	l1 := createList(list1)

	l2 := f(l1)
	l2.Print()
}

func TestInsertionSortList7(t *testing.T) {
	list1 := []int{4}
	l1 := createList(list1)

	l2 := f(l1)
	l2.Print()
}
