package leetcode

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive1(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func TestLongestConsecutive2(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, -3, 2}))
}

func TestLongestConsecutive3(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 4, -1, 1, 0, 2}))
}

func TestLongestConsecutive4(t *testing.T) {
	fmt.Println(longestConsecutive([]int{100, 99}))
}

func TestLongestConsecutive5(t *testing.T) {
	fmt.Println(longestConsecutive([]int{-2, -3}))
}
