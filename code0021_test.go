package leetcode

import "testing"

func TestMergeTwoLists1(t *testing.T) {
	list1 := []int{1, 3, 5}
	list2 := []int{2, 4, 8}

	l1 := createList(list1)
	l1.Print()

	l2 := createList(list2)
	l2.Print()

	l := mergeTwoLists(l1, l2)
	l.Print()
}

func TestMergeTwoLists2(t *testing.T) {
	list1 := []int{1, 3, 5}
	list2 := []int{2, 4, 8}

	l1 := createList(list1)
	l1.Print()

	l2 := createList(list2)
	l2.Print()

	l := mergeTwoListsRecursive(l1, l2)
	l.Print()
}
