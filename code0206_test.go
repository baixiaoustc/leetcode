package leetcode

import "testing"

func TestReverseListRecursive1(t *testing.T) {
	list1 := []int{1, 4, 3, 2, 5, 2}
	l1 := createList(list1)

	l2 := reverseListRecursive(l1)
	l2.Print()
}

func TestReverseListRecursive2(t *testing.T) {
	list1 := []int{1, 3, 5, 5, 7, 9}
	l1 := createList(list1)

	l2 := reverseListRecursive(l1)
	l2.Print()
}

func TestReverseListRecursive3(t *testing.T) {
	list1 := []int{1, 1}
	l1 := createList(list1)

	l2 := reverseListRecursive(l1)
	l2.Print()
}

func TestReverseListRecursive4(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	l2 := reverseListRecursive(l1)
	l2.Print()
}

func TestReverseListRecursive5(t *testing.T) {
	list1 := []int{}
	l1 := createList(list1)

	l2 := reverseListRecursive(l1)
	l2.Print()
}

func TestReverseListRecursive6(t *testing.T) {
	list1 := []int{1, 1, 2}
	l1 := createList(list1)

	l2 := reverseListRecursive(l1)
	l2.Print()
}

func TestReverseListRecursive7(t *testing.T) {
	list1 := []int{1, 2, 2, 0}
	l1 := createList(list1)

	l2 := reverseListRecursive(l1)
	l2.Print()
}
