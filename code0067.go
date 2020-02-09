/*
给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import "fmt"

func addBinary(a string, b string) string {
	lengtha := len(a)
	lengthb := len(b)

	aa := []rune(a)
	bb := []rune(b)
	cc := make([]rune, 0)
	carry := false
	for lengtha > 0 && lengthb > 0 {
		n := 0
		if carry {
			n++
			carry = false
		}
		n += zerone(aa[lengtha-1])
		n += zerone(bb[lengthb-1])
		if n == 2 {
			n = 0
			carry = true
		} else if n == 3 {
			n = 1
			carry = true
		}
		cc = append(cc, onezero(n))
		lengtha--
		lengthb--
	}

	for lengtha > 0 {
		n := 0
		if carry {
			n++
			carry = false
		}
		n += zerone(aa[lengtha-1])
		if n == 2 {
			n = 0
			carry = true
		}
		cc = append(cc, onezero(n))
		lengtha--
	}

	for lengthb > 0 {
		n := 0
		if carry {
			n++
			carry = false
		}
		n += zerone(bb[lengthb-1])
		if n == 2 {
			n = 0
			carry = true
		}
		cc = append(cc, onezero(n))
		lengthb--
	}

	if carry {
		cc = append(cc, '1')
	}

	fmt.Println(cc)
	dd := make([]rune, len(cc))
	for i := 0; i < len(cc); i++ {
		dd[i] = cc[len(cc)-1-i]
	}

	return string(dd)
}

func zerone(x rune) int {
	if x == '0' {
		return 0
	} else {
		return 1
	}
}

func onezero(x int) rune {
	if x == 0 {
		return '0'
	} else {
		return '1'
	}
}
