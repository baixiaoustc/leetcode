package leetcode

import (
	"fmt"
	"testing"
)

func TestCoinChangeRecursive1(t *testing.T) {
	fmt.Println(coinChangeCut([]int{1, 2, 5}, 11))
}

func TestCoinChangeRecursive2(t *testing.T) {
	fmt.Println(coinChangeCut([]int{2}, 3))
}
