/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

package leetcode

import (
	"fmt"
	"math"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	max := 0
	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			length := substring(s[i : j+1])
			if max < length {
				max = length
			}
		}
	}

	return max
}

func substring(s string) int {
	fmt.Println(s)
	list := make([]bool, 1024)
	for _, c := range s {
		if list[c] {
			return 0
		}
		list[c] = true
	}
	return len(s)
}

//滑动窗口
//https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/solution/hua-dong-chuang-kou-by-powcai/
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}

	max := 0
	left := 0
	m := make(map[int32]int)
	for i, c := range s {
		if v, ok := m[c]; ok {
			left = int(math.Max(float64(v+1), float64(left)))
		}
		m[c] = i
		max = int(math.Max(float64(i-left+1), float64(max)))
	}
	return max
}
