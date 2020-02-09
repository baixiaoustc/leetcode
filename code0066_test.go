package leetcode

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	a := []int{1, 2, 9}
	fmt.Println(plusOne1(a))
}

func TestPlusOne1(t *testing.T) {
	a := []int{9, 9, 9}
	fmt.Println(plusOne1(a))
}

func TestPlusOne2(t *testing.T) {
	a := []int{0}
	fmt.Println(plusOne1(a))
}

func TestPlusOne3(t *testing.T) {
	a := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	fmt.Println(plusOne1(a))
}

func TestSwap(t *testing.T) {
	a, b := 0, 1
	a, b = b, a
	fmt.Println(a, b)
}
