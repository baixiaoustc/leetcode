package leetcode

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	haystack := "aaaaa"
	needle := "bba"
	fmt.Println(strStrKMP(haystack, needle))
}

func TestStrStr1(t *testing.T) {
	haystack := "hello"
	needle := "ll"
	fmt.Println(strStrKMP(haystack, needle))
}

func TestStrStr2(t *testing.T) {
	haystack := "helloll"
	needle := "ll"
	fmt.Println(strStrKMP(haystack, needle))
}

func TestStrStr3(t *testing.T) {
	haystack := "lhello"
	needle := "ll"
	fmt.Println(strStrKMP(haystack, needle))
}

func TestStrStr4(t *testing.T) {
	haystack := "aaa"
	needle := "aaaa"
	fmt.Println(strStrKMP(haystack, needle))
}

func TestStrStr5(t *testing.T) {
	haystack := "aaaa"
	needle := "aaaa"
	fmt.Println(strStrKMP(haystack, needle))
}

func TestStrStr6(t *testing.T) {
	haystack := "mississipopi"
	needle := "pi"
	fmt.Println(strStrKMP(haystack, needle))
}

func TestStrStr7(t *testing.T) {
	haystack := "mississippi"
	needle := "issip"
	fmt.Println(strStrKMP(haystack, needle))
}
