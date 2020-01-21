package leetcode

import "testing"

func TestReverseBetween1(t *testing.T) {
	list1 := []int{1, 4, 3, 2, 5, 2}
	l1 := createList(list1)

	l2 := reverseBetween(l1, 2, 3)
	l2.Print()
}

func TestReverseBetween2(t *testing.T) {
	list1 := []int{1, 3, 5, 5, 7, 9}
	l1 := createList(list1)

	l2 := reverseBetween(l1, 2, 6)
	l2.Print()
}

func TestReverseBetween3(t *testing.T) {
	list1 := []int{1, 1}
	l1 := createList(list1)

	l2 := reverseBetween(l1, 1, 2)
	l2.Print()
}

func TestReverseBetween4(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	l2 := reverseBetween(l1, 1, 1)
	l2.Print()
}

func TestReverseBetween5(t *testing.T) {
	list1 := []int{}
	l1 := createList(list1)

	l2 := reverseBetween(l1, 2, 3)
	l2.Print()
}

func TestReverseBetween6(t *testing.T) {
	list1 := []int{1, 1, 2}
	l1 := createList(list1)

	l2 := reverseBetween(l1, 3, 3)
	l2.Print()
}

func TestReverseBetween7(t *testing.T) {
	list1 := []int{1, 2, 2, 0}
	l1 := createList(list1)

	l2 := reverseBetween(l1, 3, 4)
	l2.Print()
}
