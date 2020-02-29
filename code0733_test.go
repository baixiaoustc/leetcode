package leetcode

import (
	"fmt"
	"testing"
)

func TestFloodFill(t *testing.T) {
	fmt.Println(floodFill([][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}, 1, 1, 2))
}
