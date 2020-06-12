package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func Test297_0(t *testing.T) {
	obj := Constructor()

	tree := NewBinaryTree([]string{"1", "2", "3"})
	data := obj.serialize(tree)
	fmt.Println(data)
	fmt.Println(preorderTraversalRecursive(obj.deserialize(data)))
}

func Test297_1(t *testing.T) {
	obj := Constructor()

	tree := NewBinaryTree([]string{"5", "4", "7", "3", "null", "2", "null", "-1", "null", "9"})
	data := obj.serialize(tree)
	fmt.Println(data)
	fmt.Println(preorderTraversalRecursive(obj.deserialize(data)))
}

func Test297_2(t *testing.T) {
	obj := Constructor()

	tree := NewBinaryTree([]string{})
	data := obj.serialize(tree)
	fmt.Println("serialize", data, len(data))
	fmt.Println(preorderTraversalRecursive(obj.deserialize(data)))
}

func Test297(t *testing.T) {
	list := strings.Split("", ";")
	fmt.Println(list, len(list))
}
