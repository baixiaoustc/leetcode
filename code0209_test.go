package leetcode

import (
	"fmt"
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	a := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(7, a))
}

func TestMinSubArrayLen2(t *testing.T) {
	a := []int{2, 3, 1, 2, 4, 3}
	fmt.Println(minSubArrayLen(77, a))
}

func TestMinSubArrayLen3(t *testing.T) {
	a := []int{2, 3}
	fmt.Println(minSubArrayLen(5, a))
}

func TestMinSubArrayLen4(t *testing.T) {
	a := []int{2}
	fmt.Println(minSubArrayLen(2, a))
}

func TestMinSubArrayLen5(t *testing.T) {
	a := []int{}
	fmt.Println(minSubArrayLen(7, a))
}
