package leetcode

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicatesMedium1(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 4}
	count := removeDuplicatesMedium1(a1)
	fmt.Println(count, a1)
}

func TestRemoveDuplicatesMedium2(t *testing.T) {
	a1 := []int{0, 0, 1, 1, 2, 3, 4}
	count := removeDuplicatesMedium1(a1)
	fmt.Println(count, a1)
}

func TestRemoveDuplicatesMedium3(t *testing.T) {
	a1 := []int{0, 1, 2, 3, 4, 4, 4}
	count := removeDuplicatesMedium1(a1)
	fmt.Println(count, a1)
}

func TestRemoveDuplicatesMedium4(t *testing.T) {
	a1 := []int{1, 1, 1, 2, 2, 3}
	count := removeDuplicatesMedium1(a1)
	fmt.Println(count, a1)
}

func TestRemoveDuplicatesMedium5(t *testing.T) {
	a1 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	count := removeDuplicatesMedium1(a1)
	fmt.Println(count, a1)
}

func TestRemoveDuplicatesMedium6(t *testing.T) {
	a1 := []int{0, 0, 0, 0, 1, 1, 1, 1}
	count := removeDuplicatesMedium1(a1)
	fmt.Println(count, a1)
}
