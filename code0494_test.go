package leetcode

import (
	"fmt"
	"testing"
)

func TestFindTargetSumWays(t *testing.T) {
	fmt.Println(findTargetSumWaysZeroOne([]int{1, 1, 1, 1, 1}, 3))
}

func TestFindTargetSumWays1(t *testing.T) {
	fmt.Println(findTargetSumWaysZeroOne([]int{1}, 2))
}
