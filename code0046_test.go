package leetcode

import (
	"fmt"
	"testing"
)

func TestPermute1(t *testing.T) {
	fmt.Println(permute([]int{1, 2, 3}))
}

func TestPermute2(t *testing.T) {
	ret := permute([]int{5, 4, 6, 2})
	fmt.Println(len(ret), ret)
}
