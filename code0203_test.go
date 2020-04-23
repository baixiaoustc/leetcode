package leetcode

import (
	"testing"
)

func TestRemoveElements1(t *testing.T) {
	list := removeElements(createList([]int{1, 2, 6, 3, 4, 5, 6}), 6)
	list.Print()
}

func TestRemoveElements2(t *testing.T) {
	list := removeElements(createList([]int{6, 6, 6}), 6)
	list.Print()
}

func TestRemoveElements3(t *testing.T) {
	list := removeElements(createList([]int{1, 2, 3}), 6)
	list.Print()
}
