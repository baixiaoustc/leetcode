/*
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。

示例 1:

输入: nums: [1, 1, 1, 1, 1], S: 3
输出: 5
解释:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。
注意:

数组非空，且长度不会超过20。
初始的数组的和不会超过1000。
保证返回的最终结果能被32位整数存下。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/target-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

//此题很多解法，值得思考

func findTargetSumWays(nums []int, S int) int {
	count := 0
	sumDFS(nums, 0, 0, S, &count)
	return count
}

func sumDFS(nums []int, sum, index, S int, count *int) {
	if index == len(nums) {
		if sum == S {
			*count += 1
		}
		return
	}

	sumDFS(nums, sum+nums[index], index+1, S, count)
	sumDFS(nums, sum-nums[index], index+1, S, count)
}

//和上面解法类似
func findTargetSumRecursive(nums []int, S int) int {
	return sumRecursive(nums, 0, 0, S)
}

func sumRecursive(nums []int, sum, index, S int) int {
	if index == len(nums) {
		if sum == S {
			return 1
		} else {
			return 0
		}
	}

	return sumRecursive(nums, sum+nums[index], index+1, S) + sumRecursive(nums, sum-nums[index], index+1, S)
}

//剪枝
func findTargetSumWaysCut(nums []int, S int) int {
	m := make(map[[2]int]int)
	return sumDFSCUT(nums, 0, 0, S, m)
}

func sumDFSCUT(nums []int, sum, index, S int, m map[[2]int]int) int {
	if _, ok := m[[2]int{index, sum}]; !ok && (index < len(nums)) {
		m[[2]int{index, sum}] = sumDFSCUT(nums, sum+nums[index], index+1, S, m) + sumDFSCUT(nums, sum-nums[index], index+1, S, m)
	}
	if v, ok := m[[2]int{index, sum}]; ok {
		return v
	} else {
		if sum == S {
			return 1
		} else {
			return 0
		}
	}
}

//dp
func findTargetSumWaysDP(nums []int, S int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	if S > total {
		return 0
	}
	dp := make([][]int, len(nums)) //dp[i][j]，i为取到哪个数字的index，j为目前的和，值为方法数
	for i := range dp {
		dp[i] = make([]int, total*2+1)
	}

	//init
	dp[0][total+nums[0]] = 1
	dp[0][total-nums[0]] += 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= total*2; j++ {
			if j-nums[i] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
			if j+nums[i] < total*2+1 {
				dp[i][j] += dp[i-1][j+nums[i]]
			}
		}
	}
	return dp[len(nums)-1][total+S]
}

//01背包解法
func findTargetSumWaysZeroOne(nums []int, S int) int {
	var total int
	for _, n := range nums {
		total += n
	}
	if S > total {
		return 0
	}

	/* https://leetcode-cn.com/problems/target-sum/solution/python-dfs-xiang-jie-by-jimmy00745/
	                  sum(P) - sum(N) = target
	                  （两边同时加上sum(P)+sum(N)）
	sum(P) + sum(N) + sum(P) - sum(N) = target + sum(P) + sum(N)
	            (因为 sum(P) + sum(N) = sum(nums))
	                       2 * sum(P) = target + sum(nums)
	*/
	if (total+S)%2 != 0 {
		return 0
	}
	sum := (total + S) / 2

	dp := make([]int, sum+1) //dp[j]，j为目前的和，值为方法数

	//init，都是正数。取值为0的方法只有一个：什么都不取
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		for j := sum; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]

		}
	}
	return dp[sum]
}
