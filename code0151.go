/*
151. 翻转字符串里的单词
给定一个字符串，逐个翻转字符串中的每个单词。



示例 1：

输入: "the sky is blue"
输出: "blue is sky the"
示例 2：

输入: "  hello world!  "
输出: "world! hello"
解释: 输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
示例 3：

输入: "a good   example"
输出: "example good a"
解释: 如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


说明：

无空格字符构成一个单词。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。


进阶：

请选用 C 语言的用户尝试使用 O(1) 额外空间复杂度的原地解法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	ss := strings.TrimSpace(s)
	ret := make([]rune, 0)

	j := 0
	for i := len(ss) - 1; i >= 0; i-- {
		if !isSpace(rune(ss[i])) {
			j = i
		} else {
			for i > 0 && isSpace(rune(ss[i-1])) {
				i--
			}
			for j < len(ss) && !isSpace(rune(ss[j])) {
				ret = append(ret, rune(ss[j]))
				j++
			}
			ret = append(ret, rune(' '))
		}
	}

	//when i-- jump out of the ss
	for j < len(ss) && rune(ss[j]) != rune(' ') {
		ret = append(ret, rune(ss[j]))
		j++
	}

	return string(ret)
}

func isSpace(s rune) bool {
	if s == 32 || s == 160 {
		return true
	}
	return false
}

func reverseWords1(s string) string {
	start, end := 0, len(s)-1
	//先去掉头尾的空格
	for start < len(s) {
		if s[start] == ' ' {
			start++
		} else {
			break
		}
	}
	for end > start {
		if s[end] == ' ' {
			end--
		} else {
			break
		}
	}
	fmt.Println(start, end)
	ss := s[start : end+1]
	fmt.Println(ss)
	sss := reversS(ss)
	fmt.Println(sss)
	start, end = 0, 0
	var ret string
	for i := 0; i < len(sss); {
		if sss[i] != ' ' {
			end = i
			i++
		} else {
			ret += reversS(sss[start:i])
			for ; i < len(sss); i++ {
				if sss[i] != ' ' {
					ret += " "
					break
				}
			}
			fmt.Println(ret)
			start = i
		}
	}
	ret += reversS(sss[start:len(sss)])
	return ret
}

func reversS(s string) string {
	var _s string
	for i := len(s) - 1; i >= 0; i-- {
		_s += string(s[i])
	}
	return _s
}
