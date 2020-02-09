package leetcode

import (
	"fmt"
	"testing"
)

func TestFourSum1(t *testing.T) {
	a := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(a, 0))
}

func TestFourSum2(t *testing.T) {
	a := []int{-1, 0, 0, 0, 1, 2, -1, -4}
	fmt.Println(fourSum(a, 0))
}

func TestFourSum3(t *testing.T) {
	a := []int{-1, 0, 0, 0, 1, 2, -1, -4, 4, 5}
	fmt.Println(fourSum(a, 1))
}

func TestFourSum4(t *testing.T) {
	a := []int{-1, 0, 1, 2}
	fmt.Println(fourSum(a, 2))
}

func TestFourSum5(t *testing.T) {
	a := []int{0, 0, 0, 0}
	fmt.Println(fourSum(a, 0))
}

func TestFourSum6(t *testing.T) {
	a := []int{1, 1, 2}
	fmt.Println(fourSum(a, 0))
}
