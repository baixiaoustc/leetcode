package leetcode

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tree := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(preorderTraversal(tree))
}
