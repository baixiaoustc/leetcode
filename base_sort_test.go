package leetcode

import (
	"fmt"
	"testing"
)

func TestQuickSort1(t *testing.T) {
	fmt.Println(QuickSort([]int{1, 2, 3, 4, 5, 6}))
}

func TestQuickSort2(t *testing.T) {
	fmt.Println(QuickSort([]int{1, 2, 4, 4, 5, 6}))
}

func TestQuickSort3(t *testing.T) {
	fmt.Println(QuickSort([]int{6, 2, 3, 5, 4, 1}))
}

func TestQuickSort4(t *testing.T) {
	fmt.Println(QuickSort([]int{1, 1, 1, 1}))
}

func TestQuickSort5(t *testing.T) {
	fmt.Println(QuickSort([]int{6, 5, 4, 3, 2, 1}))
}
