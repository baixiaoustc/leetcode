package leetcode

import (
	"fmt"
	"testing"
)

func TestTranspose1(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(transpose(m))
}

func TestTranspose2(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}}
	fmt.Println(transpose(m))
}

func TestTranspose3(t *testing.T) {
	m := [][]int{[]int{1}, []int{4}, []int{7}}
	fmt.Println(transpose(m))
}

func TestTranspose4(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{0, 0, 0}}
	fmt.Println(transpose(m))
}
