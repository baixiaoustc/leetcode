package leetcode

import (
	"fmt"
	"testing"
)

func TestLCP(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefixDivide(strs))
}

func TestLCP1(t *testing.T) {
	strs := []string{"flower", "flow", "1flight"}
	fmt.Println(longestCommonPrefixDivide(strs))
}

func TestLCP2(t *testing.T) {
	strs := []string{"flower", "flow"}
	fmt.Println(longestCommonPrefixDivide(strs))
}

func TestLCP3(t *testing.T) {
	strs := []string{"flower"}
	fmt.Println(longestCommonPrefixDivide(strs))
}

func TestLCP4(t *testing.T) {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefixDivide(strs))
}
