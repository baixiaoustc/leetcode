/*
给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

注意:

每个数组中的元素不会超过 100
数组的大小不会超过 200
示例 1:

输入: [1, 5, 11, 5]

输出: true

解释: 数组可以分割成 [1, 5, 5] 和 [11].


示例 2:

输入: [1, 2, 3, 5]

输出: false

解释: 数组不能分割成两个元素和相等的子集.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-equal-subset-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

//普通背包问题解法
func canPartition(nums []int) bool {
	length := len(nums)
	if length == 0 || length > 200 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 || sum > 200*100 {
		return false
	}
	sum = sum / 2

	//普通背包解法
	//dp[i][j]，表示前i个组合，其和为j的可能性
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	//init
	if nums[0] <= sum {
		dp[0][nums[0]] = true
	}

	for i := 1; i < length; i++ {
		for j := 0; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j == nums[i] {
				dp[i][j] = true
			} else if j > nums[i] {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}

	return dp[length-1][sum]
}

//一维数组dp解法
//也即01背包解法（01背包的解法是内部循环为降序，对比完全背包code0322）
func canPartitionOne(nums []int) bool {
	length := len(nums)
	if length == 0 || length > 200 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 || sum > 200*100 {
		return false
	}
	sum = sum / 2

	//高级背包解法
	//dp[i]，表示和为i的可能性
	dp := make([]bool, sum+1)

	//init
	dp[0] = true

	for i := 0; i < length; i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[sum]
}
