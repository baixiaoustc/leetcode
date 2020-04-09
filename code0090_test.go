package leetcode

import (
	"fmt"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func TestSubsetsWithDup1(t *testing.T) {
	fmt.Println(subsetsWithDup([]int{1, 4, 4, 4, 4}))
}
