package leetcode

import (
	"fmt"
	"testing"
)

func TestThreeSum1(t *testing.T) {
	a := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))
}

func TestThreeSum2(t *testing.T) {
	a := []int{-1, 0, 0, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(a))
}

func TestThreeSum3(t *testing.T) {
	a := []int{-1, 0, 0, 0, 1, 2, -1, -4, 4, 5}
	fmt.Println(threeSum(a))
}

func TestThreeSum4(t *testing.T) {
	a := []int{-1, 0, 1}
	fmt.Println(threeSum(a))
}

func TestThreeSum5(t *testing.T) {
	a := []int{0, 0, 0}
	fmt.Println(threeSum(a))
}

func TestThreeSum6(t *testing.T) {
	a := []int{1, 1, 2}
	fmt.Println(threeSum(a))
}
