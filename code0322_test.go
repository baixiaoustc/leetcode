package leetcode

import (
	"fmt"
	"testing"
)

func TestCoinChangeRecursive1(t *testing.T) {
	fmt.Println(coinChangeDP([]int{1, 2, 5}, 11))
}

func TestCoinChangeRecursive2(t *testing.T) {
	fmt.Println(coinChangeDP([]int{2}, 3))
}
