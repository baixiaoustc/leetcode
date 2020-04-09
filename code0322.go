/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
说明:
你可以认为每种硬币的数量是无限的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/coin-change
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import "math"

//普通递归 OOT (out of time)
func coinChangeRecursive(coins []int, amount int) int {
	return exchangeRecursive(coins, amount)
}

//递归函数，返回凑硬币的最小个数，不能凑的情况返回-1
func exchangeRecursive(coins []int, amount int) int {
	//两个临界条件
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}

	//因为是求最小值，所以先初始化为极大值
	var result int = math.MaxInt32
	for _, coin := range coins {
		ret := exchangeRecursive(coins, amount-coin)
		if ret == -1 {
			continue
		}
		result = int(math.Min(float64(result), float64(ret)+1))
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}

func coinChangeCut(coins []int, amount int) int {
	m := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		m[i] = math.MaxInt32
	}
	return exchangeCut(coins, amount, m)
}

//递归函数，加剪枝
func exchangeCut(coins []int, amount int, m []int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	//存储中间结果
	if m[amount] != math.MaxInt32 {
		return m[amount]
	}

	var result int = math.MaxInt32
	for _, coin := range coins {
		ret := exchangeCut(coins, amount-coin, m)
		if ret == -1 {
			continue
		}
		result = int(math.Min(float64(result), float64(ret)+1))
	}

	if result == math.MaxInt32 {
		result = -1
	}
	m[amount] = result

	return result
}

//可以从递推公式的角度来理解
//也可以从完全背包解法来记忆（完全背包解法是内部循环为升序，对比01背包code0416）
func coinChangeDP(coins []int, amount int) int {
	dp := make([]int, amount+1) //dp[i]表示amount为i时的最优解
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}

	//init
	dp[0] = 0

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			//递推公式
			//f(amount, coins) = min{ f(amount - coins[i]) + 1) }, 其中 i 的取值为 0 到 coins 的大小，coins[i] 表示选择了此硬币, 1 表示选择了coins[i]  这一枚硬币
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}
