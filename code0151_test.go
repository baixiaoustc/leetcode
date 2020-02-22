package leetcode

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	fmt.Println(reverseWords557("the sky is blue"))
}

func TestReverseWords1(t *testing.T) {
	fmt.Println(reverseWords557("  hello world!  "))
}

func TestReverseWords2(t *testing.T) {
	fmt.Println(reverseWords557("a good   example"))
}

func TestReverseWords3(t *testing.T) {
	fmt.Println(reverseWords557("Let's take LeetCode contest"))
}

func TestRune1(t *testing.T) {
	s := "   "
	fmt.Println(s, len(s))
	for i, n := range s {
		fmt.Println(i, n)
	}
}
