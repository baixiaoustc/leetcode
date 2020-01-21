package leetcode

import (
	"fmt"
	"testing"
)

func TestFindList(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	n0 := findList(l1, 0)
	fmt.Println(n0.String())

	n1 := findList(l1, 1)
	fmt.Println(n1.String())

	n5 := findList(l1, 5)
	fmt.Println(n5.String())

	n8 := findList(l1, 8)
	fmt.Println(n8.String())
}

func TestCutList(t *testing.T) {
	list1 := []int{1, 3, 5}
	l1 := createList(list1)

	l2 := cutList(l1, 1)
	l2.Print()

	list3 := []int{1, 3, 5}
	l3 := createList(list3)

	l4 := cutList(l3, 0)
	l4.Print()

	list5 := []int{1, 3, 5}
	l5 := createList(list5)

	l6 := cutList(l5, 3)
	l6.Print()
}

func TestReverseList(t *testing.T) {
	list1 := []int{1, 3, 5}
	l1 := createList(list1)
	l2 := reverseList(l1)
	l2.Print()

	list3 := []int{1, 3}
	l3 := createList(list3)
	l4 := reverseList(l3)
	l4.Print()

	list5 := []int{1}
	l5 := createList(list5)
	l6 := reverseList(l5)
	l6.Print()
}
