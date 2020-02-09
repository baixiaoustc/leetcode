/*
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
在真实的面试中遇到过这道

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/4sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	ret := make([][]int, 0)
	length := len(nums)
	if length < 4 {
		return ret
	}

	sort.Ints(nums)

	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue //skip duplicate
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue //skip duplicate
			}

			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			if nums[length-1]+nums[length-2]+nums[length-3]+nums[length-4] < target {
				break
			}

			l := j + 1
			r := length - 1
			for {
				if l >= r {
					break
				}

				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--
					for nums[l] == nums[l-1] && l <= r {
						l++
					}
					for nums[r] == nums[r+1] && l <= r {
						r--
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
		}
	}

	return ret
}
