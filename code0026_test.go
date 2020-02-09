package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates1(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 4}
	count := removeDuplicates(a1)
	fmt.Println(count, a1[0:count])
}

func TestRemoveDuplicates2(t *testing.T) {
	a1 := []int{0, 0, 1, 1, 2, 3, 4}
	count := removeDuplicates(a1)
	fmt.Println(count, a1[0:count])
}

func TestRemoveDuplicates3(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 4, 4, 4}
	count := removeDuplicates(a1)
	fmt.Println(count, a1[0:count])
}

func TestRemoveDuplicates4(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 3, 5, 5}
	count := removeDuplicates(a1)
	fmt.Println(count, a1[0:count])
}

func TestRemoveDuplicates5(t *testing.T) {
	a1 := []int{}
	count := removeDuplicates(a1)
	fmt.Println(count, a1[0:count])
}
