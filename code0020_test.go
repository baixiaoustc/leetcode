package leetcode

import (
	"fmt"
	"testing"
)

func TestIsValid1(t *testing.T) {
	fmt.Println(isValid("()"))
}

func TestIsValid2(t *testing.T) {
	fmt.Println(isValid("()[]{}"))
}

func TestIsValid3(t *testing.T) {
	fmt.Println(isValid("(]"))
}

func TestIsValid4(t *testing.T) {
	fmt.Println(isValid("([)]"))
}

func TestIsValid5(t *testing.T) {
	fmt.Println(isValid("{[]}"))
}
