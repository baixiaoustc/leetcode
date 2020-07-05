/*
567. 字符串的排列
给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。

换句话说，第一个字符串的排列之一是第二个字符串的子串。

示例1:

输入: s1 = "ab" s2 = "eidbaooo"
输出: True
解释: s2 包含 s1 的排列之一 ("ba").


示例2:

输入: s1= "ab" s2 = "eidboaoo"
输出: False


注意：

输入的字符串只包含小写字母
两个字符串的长度都在 [1, 10,000] 之间
*/

package leetcode

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	lens1 := len(s1)
	ms1 := make(map[string]int)
	for _, c := range s1 {
		ms1[string(c)] = ms1[string(c)] + 1
	}
	fmt.Println(s1, ms1)

	left := 0
	ms2 := make(map[string]int)
	for i, c := range s2 {
		//fmt.Println(string(c))
		if _, ok := ms1[string(c)]; !ok {
			fmt.Println(string(c), "not in")
			left = i + 1
			ms2 = make(map[string]int)
			continue
		}
		ms2[string(c)] = ms2[string(c)] + 1
		//fmt.Println(left, ms2)
		if i-left+1 == lens1 { //the same length, compare
			var same bool = true
			for k, v := range ms1 {
				if v != ms2[k] {
					same = false
					break
				}
			}
			if same {
				return true
			} else {
				fmt.Println(left, s2[left], string(s2[left]), s2[left:i], "not equal")
				fmt.Println(ms2)
				ms2[string(s2[left])] = ms2[string(s2[left])] - 1
				left++
			}
		}
	}
	return false
}
