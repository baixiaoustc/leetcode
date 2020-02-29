package leetcode

import (
	"fmt"
	"testing"
)

func TestDecodeString1(t *testing.T) {
	fmt.Println(decodeString("3[a]2[bc]"))
}

func TestDecodeString2(t *testing.T) {
	fmt.Println(decodeString("3[a2[c]]"))
}

func TestDecodeString3(t *testing.T) {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}

func TestDecodeString4(t *testing.T) {
	fmt.Println(decodeString("100[leetcode]"))
}
