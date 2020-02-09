/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

//分治法
func longestCommonPrefixDivide(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	return lcpDivide(strs, 0, len(strs)-1)
}

func lcpDivide(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}

	mid := (l + r) / 2
	left := lcpDivide(strs, l, mid)
	right := lcpDivide(strs, mid+1, r)

	length := len(left)
	if length > len(right) {
		length = len(right)
	}

	for i := 0; i < length; i++ {
		if left[i:i+1] != right[i:i+1] {
			return left[0:i]
		}
	}

	return left[0:length]
}
