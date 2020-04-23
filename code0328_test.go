package leetcode

import "testing"

func TestOddEvenList1(t *testing.T) {
	list := oddEvenList(createList([]int{1, 2, 3, 4, 5}))
	list.Print()
}

func TestOddEvenList2(t *testing.T) {
	list := oddEvenList(createList([]int{2, 1, 3, 5, 6, 4, 7}))
	list.Print()
}

func TestOddEvenList3(t *testing.T) {
	list := oddEvenList(createList([]int{1, 2, 3, 4, 5, 6, 7, 8}))
	list.Print()
}

func TestOddEvenList4(t *testing.T) {
	list := oddEvenList(createList([]int{1}))
	list.Print()
}

func TestOddEvenList5(t *testing.T) {
	list := oddEvenList(createList([]int{1, 2, 3, 4}))
	list.Print()
}
