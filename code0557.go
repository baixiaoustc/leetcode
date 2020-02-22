/*
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc"
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func reverseWords557(s string) string {
	if s == "" {
		return s
	}

	ret := make([]rune, 0)
	j := 0
	for i := 0; i < len(s); i++ {
		//if isSpace(rune(s[i])) {
		if s[i:i+1] == " " {
			//find space
			for k := i - 1; k >= j; k-- {
				ret = append(ret, rune(s[k]))
			}
			ret = append(ret, rune(' '))
			j = i + 1
		}
	}

	for k := len(s) - 1; k >= j; k-- {
		ret = append(ret, rune(s[k]))
	}

	return string(ret)
}
