package leetcode

import (
	"fmt"
	"testing"
)

func TestPivotIndex1(t *testing.T) {
	a := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(a))
}

func TestPivotIndex2(t *testing.T) {
	a := []int{1, 2, 3}
	fmt.Println(pivotIndex(a))
}

func TestPivotIndex3(t *testing.T) {
	a := []int{0, 0, 0}
	fmt.Println(pivotIndex(a))
}

func TestPivotIndex4(t *testing.T) {
	a := []int{1, -1, 0}
	fmt.Println(pivotIndex(a))
}

func TestHi(t *testing.T) {
	s := Student{}
	s.Init("pirlo", 21, "cs")
	s.SayHi()
}
