package leetcode

import "testing"

func TestRemoveNthFromEnd1(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := removeNthFromEndOneScan1(l1, 2)
	l2.Print()
}

func TestRemoveNthFromEnd2(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := removeNthFromEndOneScan1(l1, 0)
	l2.Print()
}

func TestRemoveNthFromEnd4(t *testing.T) {
	list1 := []int{1, 3}
	l1 := createList(list1)

	l2 := removeNthFromEndOneScan1(l1, 1)
	l2.Print()
}

func TestRemoveNthFromEnd5(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	l2 := removeNthFromEndOneScan1(l1, 1)
	l2.Print()
}

func TestRemoveNthFromEnd6(t *testing.T) {
	list1 := []int{1, 2}
	l1 := createList(list1)

	l2 := removeNthFromEndOneScan1(l1, 2)
	l2.Print()
}
