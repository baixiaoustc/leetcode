package leetcode

import "testing"

func TestDeleteDuplicates1(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := deleteDuplicates83(l1)
	l2.Print()
}

func TestDeleteDuplicates2(t *testing.T) {
	list1 := []int{1, 3, 5, 5, 7, 9}
	l1 := createList(list1)

	l2 := deleteDuplicates83(l1)
	l2.Print()
}

func TestDeleteDuplicates3(t *testing.T) {
	list1 := []int{1, 1}
	l1 := createList(list1)

	l2 := deleteDuplicates83(l1)
	l2.Print()
}

func TestDeleteDuplicates4(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	l2 := deleteDuplicates83(l1)
	l2.Print()
}

func TestDeleteDuplicates5(t *testing.T) {
	list1 := []int{}
	l1 := createList(list1)

	l2 := deleteDuplicates83(l1)
	l2.Print()
}

func TestDeleteDuplicates6(t *testing.T) {
	list1 := []int{1, 1, 2}
	l1 := createList(list1)

	l2 := deleteDuplicates83(l1)
	l2.Print()
}

func TestDeleteDuplicates7(t *testing.T) {
	list1 := []int{1, 2, 2, 2}
	l1 := createList(list1)

	l2 := deleteDuplicates83(l1)
	l2.Print()
}
