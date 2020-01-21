package leetcode

import "testing"

func TestSwapPairs1(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := swapPairs(l1)
	l2.Print()
}

func TestSwapPairs2(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9}
	l1 := createList(list1)

	l2 := swapPairs(l1)
	l2.Print()
}

func TestSwapPairs3(t *testing.T) {
	list1 := []int{1, 3}
	l1 := createList(list1)

	l2 := swapPairs(l1)
	l2.Print()
}

func TestSwapPairs4(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	l2 := swapPairs(l1)
	l2.Print()
}

func TestSwapPairs5(t *testing.T) {
	list1 := []int{}
	l1 := createList(list1)

	l2 := swapPairs(l1)
	l2.Print()
}
