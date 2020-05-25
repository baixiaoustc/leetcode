package leetcode

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	tree := buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3})
	fmt.Println(preorderTraversal(tree))
}

func TestBuildTree0105(t *testing.T) {
	tree := buildTree0105([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(preorderTraversal(tree))
}
