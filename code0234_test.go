package leetcode

import (
	"fmt"
	"testing"
)

func TestIsPalindrome1(t *testing.T) {
	fmt.Println(isPalindrome(createList([]int{1, 2})))
}

func TestIsPalindrome2(t *testing.T) {
	fmt.Println(isPalindrome(createList([]int{1, 2, 2, 1})))
}

func TestIsPalindrome3(t *testing.T) {
	fmt.Println(isPalindrome(createList([]int{1, 2, 3, 2, 1})))
}

func TestIsPalindrome4(t *testing.T) {
	fmt.Println(isPalindrome(createList([]int{1, 2, 3, 2, 2})))
}
