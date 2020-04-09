/*
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import "math"

func numSquares(n int) int {
	if n == 0 {
		return 0
	}

	visit := make([]int, n+1)

	queue := NewLinkedQueue()
	queue.Add([2]int{n, 0})
	for queue.Size() != 0 {
		ele, err := queue.Peek()
		if err != nil {
			panic("queue peek")
		}
		err = queue.Remove()
		if err != nil {
			panic("queue remove")
		}

		next := ele.([2]int)[0]
		step := ele.([2]int)[1]

		step++
		for i := 1; ; i++ {
			if del := next - i*i; del < 0 {
				break
			} else if del == 0 {
				return step
			} else if visit[del] == 0 {
				queue.Add([2]int{del, step})
				visit[del] = 1
			}
		}
	}

	return -1
}

func numSquaresZeroOne(n int) int {
	squareNums := make([]int, 0)
	for i := 1; ; i++ {
		if i*i > n {
			break
		}
		squareNums = append(squareNums, i*i)
	}

	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	//init
	dp[0] = 0

	for _, square := range squareNums {
		for i := square; i <= n; i++ {
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-square]+1)))
		}
	}

	if dp[n] == math.MaxInt32 {
		return -1
	}
	return dp[n]
}
