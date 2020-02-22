package leetcode

import (
	"fmt"
	"testing"
)

func TestEvalRPN1(t *testing.T) {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}

func TestEvalRPN2(t *testing.T) {
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
}

func TestEvalRPN3(t *testing.T) {
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func TestEvalRPN4(t *testing.T) {
	fmt.Println(evalRPN([]string{"0", "0", "+", "0", "*"}))
}
