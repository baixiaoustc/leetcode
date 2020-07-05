/*
43. 字符串相乘
给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例 2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1 和 num2 的长度小于110。
num1 和 num2 只包含数字 0-9。
num1 和 num2 均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

package leetcode

import "fmt"

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	ret := make([]uint8, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			p0, p1, p2 := i+j-1, i+j, i+j+1
			mul := (num1[i] - '0') * (num2[j] - '0')
			fmt.Println(mul)
			ret[p2] += mul % 10
			if ret[p2] >= 10 {
				ret[p1]++
				ret[p2] -= 10
			}
			ret[p1] += mul / 10
			if ret[p1] >= 10 {
				ret[p0]++
				ret[p1] -= 10
			}
		}
	}
	fmt.Println(ret)
	var s string
	var i int
	for ; i < len(ret); i++ {
		if ret[i] == 0 {
			continue
		} else {
			break
		}
	}
	fmt.Println(i)
	for ; i < len(ret); i++ {
		s += fmt.Sprint(ret[i])
	}
	return s
}
