/*
121. 买卖股票的最佳时机
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。



示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
*/

package leetcode

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([]int, 2) //两种状态
	}

	//init
	dp[0][0] = 0
	dp[0][1] = -math.MaxInt32 //还没开始的时候不可能有股票，用负无穷表示不可能

	for i := 1; i < len(prices)+1; i++ {
		dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i-1])))
		dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(-prices[i-1]))) //这个地方notice
	}

	//for i := 0; i < len(prices); i++ {
	//	if i == 0 {
	//		dp[i][0] = 0
	//		dp[i][1] = -math.MaxInt32
	//		continue
	//	}
	//	dp[i][0] = int(math.Max(float64(dp[i-1][0]), float64(dp[i-1][1]+prices[i])))
	//	dp[i][1] = int(math.Max(float64(dp[i-1][1]), float64(-prices[i])))
	//}

	return dp[len(prices)][0]
}

func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	dp0, dp1 := 0, -math.MaxInt32
	for i := 0; i < len(prices); i++ {
		dp0 = int(math.Max(float64(dp0), float64(dp1+prices[i])))
		dp1 = int(math.Max(float64(dp1), float64(-prices[i])))
	}

	return dp0
}
