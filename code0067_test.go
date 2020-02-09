package leetcode

import (
	"fmt"
	"testing"
)

func TestAddBinary1(t *testing.T) {
	a := "11"
	b := "1"
	fmt.Println(addBinary(a, b))
}

func TestAddBinary2(t *testing.T) {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b))
}

func TestAddBinary3(t *testing.T) {
	a := "0"
	b := "0"
	fmt.Println(addBinary(a, b))
}

func TestAddBinary4(t *testing.T) {
	a := "11"
	b := "11"
	fmt.Println(addBinary(a, b))
}
