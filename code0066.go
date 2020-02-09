/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

//可能导致溢出
func plusOne(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{}
	}

	sum := 0
	max := 0
	for _, n := range digits {
		sum = sum*10 + n
		max = max*10 + 9
	}

	if sum == max {
		length++
	}
	sum++

	tmp := 0
	ret := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		tmp = sum / 10
		ret[i] = sum - tmp*10
		sum = tmp
	}
	return ret
}

func plusOne1(digits []int) []int {
	length := len(digits)
	if length == 0 {
		return []int{}
	}

	var carry bool
	ret := make([]int, length)
	for i := length - 1; i >= 0; i-- {
		n := digits[i]
		if i == length-1 {
			n++
		}
		if carry {
			n++
			carry = false
		}
		if n == 10 {
			n = 0
			carry = true
		}
		ret[i] = n
	}

	if carry {
		ret = make([]int, length+1)
		ret[0] = 1
	}

	return ret
}
