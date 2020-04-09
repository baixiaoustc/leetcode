package leetcode

import (
	"testing"
	"fmt"
)

func TestLongestCommonSubsequence1(t *testing.T)  {
	fmt.Println(longestCommonSubsequence("abcde","ace"))
}

func TestLongestCommonSubsequence2(t *testing.T)  {
	fmt.Println(longestCommonSubsequence("abc","abc"))
}

func TestLongestCommonSubsequence3(t *testing.T)  {
	fmt.Println(longestCommonSubsequence("abc","def"))
}
