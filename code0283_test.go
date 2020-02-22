package leetcode

import (
	"fmt"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	a := []int{0, 1, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}

func TestMoveZeros1(t *testing.T) {
	a := []int{0, 0, 0, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}

func TestMoveZeros2(t *testing.T) {
	a := []int{0, 1, 0, 3, 12, 0}
	moveZeroes(a)
	fmt.Println(a)
}

func TestMoveZeros3(t *testing.T) {
	a := []int{0, 0}
	moveZeroes(a)
	fmt.Println(a)
}

func TestMoveZeros4(t *testing.T) {
	a := []int{1, 3, 12}
	moveZeroes(a)
	fmt.Println(a)
}
