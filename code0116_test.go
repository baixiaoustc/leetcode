package leetcode

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	tree := NewPerfectTree([]string{"1", "2", "3", "4", "5", "6", "7"})
	fmt.Println(connect(tree))
}
