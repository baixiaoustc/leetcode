package leetcode

import (
	"fmt"
	"testing"
)

func TestNewBinaryTree1(t *testing.T) {
	tree := NewBinaryTree([]string{"1", "2", "3"})
	fmt.Println(preorderTraversalRecursive(tree))
}

func TestNewBinaryTree2(t *testing.T) {
	tree := NewBinaryTree([]string{"1", "2", "3", "4", "5"})
	fmt.Println(preorderTraversalRecursive(tree))
}

func TestNewBinaryTree3(t *testing.T) {
	tree := NewBinaryTree([]string{"1", "2", "3", "4", "null", "5"})
	fmt.Println(preorderTraversalRecursive(tree))
}

func TestNewBinaryTree4(t *testing.T) {
	tree := NewBinaryTree([]string{"1", "2", "null", "4", "5"})
	fmt.Println(preorderTraversalRecursive(tree))
}

func TestNewBinaryTree5(t *testing.T) {
	tree := NewBinaryTree([]string{"1", "2", "2", "2", "null", "2"})
	fmt.Println(preorderTraversalRecursive(tree))
}
