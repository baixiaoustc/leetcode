package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveElement1(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 4}
	fmt.Println(removeElement(a1, 6))
}

func TestRemoveElement2(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 4, 0}
	fmt.Println(removeElement(a1, 0))
}

func TestRemoveElement3(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 0, 0, 4}
	fmt.Println(removeElement(a1, 0))
}
