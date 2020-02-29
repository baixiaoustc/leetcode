/*
给定一个经过编码的字符串，返回它解码后的字符串。

编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例:

s = "3[a]2[bc]", 返回 "aaabcbc".
s = "3[a2[c]]", 返回 "accaccacc".
s = "2[abc]3[cd]ef", 返回 "abcabccdcdcdef".

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	if s == "" {
		return ""
	}

	var ret string
	var x int
	stack := NewStackInterface()
	for _, n := range s {
		if n >= '0' && n <= '9' {
			_n, _ := strconv.Atoi(string(n))
			x = x*10 + _n
		} else if n == '[' {
			fmt.Println("in", x)
			stack.Push(x)
			x = 0
		} else if n == ']' {
			var c string
			var y int
		NOOP:
			for stack.Size() != 0 {
				_c := stack.Pop()
				//fmt.Println(_c)
				switch _c.(type) {
				case int:
					y = _c.(int)
					break NOOP
				case string:
					c = fmt.Sprintf("%s%s", c, _c)
				}
			}
			fmt.Println("out", y, c)
			for i := 0; i < y; i++ {
				fmt.Println("in", c)
				stack.Push(c)
			}
		} else {
			stack.Push(string(n))
		}
	}

	for stack.Size() != 0 {
		_c := stack.Pop().(string)
		var c []rune = make([]rune, len(_c))
		for i, n := range _c {
			c[len(c)-1-i] = n
		}
		ret = fmt.Sprintf("%s%s", string(c), ret)
	}

	return ret
}
