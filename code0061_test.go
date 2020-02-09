package leetcode

import (
	"fmt"
	"testing"
)

func TestRotateRight1(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := rotateRight(l1, 1)
	l2.Print()
}

func TestRotateRight2(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := rotateRight(l1, 12)
	l2.Print()
}

func TestRotateRight3(t *testing.T) {
	list1 := []int{1, 3, 5, 7, 9, 11}
	l1 := createList(list1)

	l2 := rotateRight(l1, 5)
	l2.Print()
}

func TestRotateRight4(t *testing.T) {
	list1 := []int{1, 3}
	l1 := createList(list1)

	l2 := rotateRight(l1, 1)
	l2.Print()
}

func TestRotateRight5(t *testing.T) {
	list1 := []int{1, 3}
	l1 := createList(list1)

	l2 := rotateRight(l1, 3)
	l2.Print()
}

func TestRotateRight6(t *testing.T) {
	list1 := []int{1}
	l1 := createList(list1)

	l2 := rotateRight(l1, 1)
	l2.Print()
}

func TestRotate1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	fmt.Println(a)
}

func TestRotate2(t *testing.T) {
	a := []int{1, 2}
	rotate(a, 3)
	fmt.Println(a)
}

func TestRotate3(t *testing.T) {
	a := []int{1}
	rotate(a, 3)
	fmt.Println(a)
}
