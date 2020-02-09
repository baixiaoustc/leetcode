/*
给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。

示例 1:

输入: [1,4,3,2]

输出: 4
解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
提示:

n 是正整数,范围在 [1, 10000].
数组中的元素范围在 [-10000, 10000].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/array-partition-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import (
	"math"
	"sort"
)

func arrayPairSum(nums []int) int {
	if len(nums)%2 != 0 {
		return 0
	}

	sort.Ints(nums)

	j := 0
	sum := 0
	for i := 1; i < len(nums); i += 2 {
		sum += int(math.Min(float64(nums[i]), float64(nums[j])))
		j += 2
	}

	return sum
}
