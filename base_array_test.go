package leetcode

import (
	"fmt"
	"testing"
)

func TestDelArray1(t *testing.T) {
	a1 := []int{0, 1, 2}
	fmt.Println(delArray(a1, 0))
}

func TestDelArray2(t *testing.T) {
	a1 := []int{0, 1, 2}
	fmt.Println(delArray(a1, 2))
}

func TestDelArray3(t *testing.T) {
	a1 := []int{0}
	fmt.Println(delArray(a1, 0))
}

func TestDelArray(t *testing.T) {
	a1 := []int{}
	fmt.Println(delArray(a1, 0))
}
