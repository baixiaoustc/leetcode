package leetcode

import "testing"

func TestAddTwoNumbers1(t *testing.T) {
	list1 := []int{4}
	l1 := createList(list1)

	list2 := []int{3}
	l2 := createList(list2)

	l := addTwoNumbers(l1, l2)
	l.Print()
}

func TestAddTwoNumbers2(t *testing.T) {
	list1 := []int{4}
	l1 := createList(list1)

	list2 := []int{9}
	l2 := createList(list2)

	l := addTwoNumbers(l1, l2)
	l.Print()
}

func TestAddTwoNumbers3(t *testing.T) {
	list1 := []int{4, 9}
	l1 := createList(list1)

	list2 := []int{9}
	l2 := createList(list2)

	l := addTwoNumbers(l1, l2)
	l.Print()
}

func TestAddTwoNumbers4(t *testing.T) {
	list1 := []int{4, 9}
	l1 := createList(list1)

	list2 := []int{9, 0}
	l2 := createList(list2)

	l := addTwoNumbers(l1, l2)
	l.Print()
}

func TestAddTwoNumbers5(t *testing.T) {
	list1 := []int{4}
	l1 := createList(list1)

	list2 := []int{9, 0, 9}
	l2 := createList(list2)

	l := addTwoNumbers(l1, l2)
	l.Print()
}

func TestAddTwoNumbers6(t *testing.T) {
	list1 := []int{0}
	l1 := createList(list1)

	list2 := []int{9, 0}
	l2 := createList(list2)

	l := addTwoNumbers(l1, l2)
	l.Print()
}

func TestAddTwoNumbers7(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	list2 := []int{9, 9}
	l2 := createList(list2)

	l := addTwoNumbers(l1, l2)
	l.Print()
}
