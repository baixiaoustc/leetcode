package leetcode

import "testing"

func TestPartition1(t *testing.T) {
	list1 := []int{1, 4, 3, 2, 5, 2}
	l1 := createList(list1)

	l2 := partition(l1, 3)
	l2.Print()
}

func TestPartition2(t *testing.T) {
	list1 := []int{1, 3, 5, 5, 7, 9}
	l1 := createList(list1)

	l2 := partition(l1, 5)
	l2.Print()
}

func TestPartition3(t *testing.T) {
	list1 := []int{1, 1}
	l1 := createList(list1)

	l2 := partition(l1, 3)
	l2.Print()
}

func TestPartition4(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	l2 := partition(l1, 3)
	l2.Print()
}

func TestPartition5(t *testing.T) {
	list1 := []int{}
	l1 := createList(list1)

	l2 := partition(l1, 3)
	l2.Print()
}

func TestPartition6(t *testing.T) {
	list1 := []int{1, 1, 2}
	l1 := createList(list1)

	l2 := partition(l1, 1)
	l2.Print()
}

func TestPartition7(t *testing.T) {
	list1 := []int{1, 2, 2, 0}
	l1 := createList(list1)

	l2 := partition(l1, 2)
	l2.Print()
}
