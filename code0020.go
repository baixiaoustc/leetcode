/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func isValid(s string) bool {
	left := func(s string) bool {
		if s == "(" || s == "[" || s == "{" {
			return true
		}
		return false
	}

	match := func(s1, s2 string) bool {
		if s1 == "(" && s2 == ")" {
			return true
		}
		if s1 == "[" && s2 == "]" {
			return true
		}
		if s1 == "{" && s2 == "}" {
			return true
		}
		return false
	}

	stack := NewStackInterface()
	for _, c := range s {
		c1 := string(c)
		if left(c1) {
			stack.Push(c1)
		} else {
			if stack.Size() == 0 {
				return false
			}
			c2 := stack.Pop().(string)
			if !match(c2, c1) {
				return false
			}
		}
	}

	if stack.Size() != 0 {
		return false
	}

	return true
}
