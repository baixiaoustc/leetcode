package leetcode

import (
	"fmt"
	"testing"
)

func TestMultiply1(t *testing.T) {
	fmt.Println(multiply("1", "9"))
}

func TestMultiply2(t *testing.T) {
	fmt.Println(multiply("17", "99"), 17*99)
}

func TestMultiply3(t *testing.T) {
	fmt.Println(multiply("172", "12"), 172*12)
}

func TestMultiply4(t *testing.T) {
	fmt.Println(multiply("872", "92"), 872*92)
}
