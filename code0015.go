/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	length := len(nums)
	if length < 3 {
		return ret
	}

	sort.Ints(nums)

	for i, n := range nums {
		if n > 0 {
			break
		}
		if i > 0 && n == nums[i-1] {
			continue //skip duplicate
		}

		l := i + 1
		r := length - 1
		for {
			if l >= r {
				break
			}

			sum := n + nums[l] + nums[r]
			if sum == 0 {
				ret = append(ret, []int{n, nums[l], nums[r]})
				l++
				r--
				for nums[l] == nums[l-1] && l <= r {
					l++
				}
				for nums[r] == nums[r+1] && l <= r {
					r--
				}
			} else if sum < 0 {
				l++
			} else {
				r--
			}
		}
	}

	return ret
}
