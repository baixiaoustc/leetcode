/*
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。

我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。

如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

示例 1:

输入:
nums = [1, 7, 3, 6, 5, 6]
输出: 3
解释:
索引3 (nums[3] = 6) 的左侧数之和(1 + 7 + 3 = 11)，与右侧数之和(5 + 6 = 11)相等。
同时, 3 也是第一个符合要求的中心索引。
示例 2:

输入:
nums = [1, 2, 3]
输出: -1
解释:
数组中不存在满足此条件的中心索引。
说明:

nums 的长度范围为 [0, 10000]。
任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-pivot-index
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import "fmt"

func pivotIndex(nums []int) int {
	tmp, sum1, sum2 := 0, 0, 0
	for i, n := range nums {

		if i == 0 {
			for j := 1; j < len(nums); j++ {
				sum2 += nums[j]
			}
			if sum2 == 0 {
				return 0
			}
			tmp = n
			continue
		}

		sum1 += tmp
		sum2 -= n
		if sum1 == sum2 {
			return i
		}

		tmp = n
	}
	return -1
}

type Student struct {
	name  string
	age   int
	major string
}

func (s *Student) Init(name string, age int, major string) {
	s.name = name
	s.age = age
	s.major = major
}

func (s Student) SayHi() {
	fmt.Printf("Hi, I am %s aged %d, and my major is %s\n", s.name, s.age, s.major)
}
