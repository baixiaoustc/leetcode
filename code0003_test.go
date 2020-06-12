package leetcode

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring1(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring1("abcabcbb"))
}

func TestLengthOfLongestSubstring2(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring1("bbbbb"))
}

func TestLengthOfLongestSubstring3(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring1("pwwkew"))
}

func TestLengthOfLongestSubstring4(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring1(" "))
}

func TestLengthOfLongestSubstring5(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring1("au"))
}
