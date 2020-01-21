package leetcode

import "testing"

func TestSortList1(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := sortList(l1)
	l2.Print()
}

func TestSortList2(t *testing.T) {
	list1 := []int{1, 3, 2, 7, 6, 11}
	l1 := createList(list1)

	l2 := sortList(l1)
	l2.Print()
}

func TestSortList3(t *testing.T) {
	list1 := []int{1, 3, 3, 7, 9, 11, 9}
	l1 := createList(list1)

	l2 := sortList(l1)
	l2.Print()
}

func TestSortList4(t *testing.T) {
	list1 := []int{1, 3, -5, 7, 0, 11}
	l1 := createList(list1)

	l2 := sortList(l1)
	l2.Print()
}
