package leetcode

import (
	"fmt"
	"testing"
)

func TestFindDiagonalOrder1(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(findDiagonalOrder(m))
}

func TestFindDiagonalOrder2(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}}
	fmt.Println(findDiagonalOrder(m))
}

func TestFindDiagonalOrder3(t *testing.T) {
	m := [][]int{[]int{1}, []int{4}, []int{7}}
	fmt.Println(findDiagonalOrder(m))
}

func TestFindDiagonalOrder4(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{0, 0, 0}}
	fmt.Println(findDiagonalOrder(m))
}
