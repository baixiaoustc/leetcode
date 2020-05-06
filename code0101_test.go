package leetcode

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsSymmetric1(t *testing.T) {
	tree := NewBinaryTree([]string{"1", "2", "2", "2", "null", "2"})
	fmt.Println(isSymmetric(tree))
}

func TestIsSymmetric2(t *testing.T) {
	tree := NewBinaryTree([]string{"4", "-57", "-57", "null", "67", "67", "null", "null", "-97", "-97"})
	fmt.Println(isSymmetric(tree))

	s := "@1.-57@@2.67@@3.-97@"
	fmt.Println(strings.Split(s, "@"))
}
