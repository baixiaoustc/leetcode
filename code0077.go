/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func combine(n int, k int) [][]int {
	path := make([]int, 0)
	ret := make([][]int, 0)
	backtrackCombine(n, k, path, &ret, 1)
	return ret
}

func backtrackCombine(n, k int, path []int, ret *[][]int, start int) {
	if k == len(path) {
		_path := make([]int, len(path))
		copy(_path, path)
		*ret = append(*ret, _path)
		return
	}

	for i := start; i <= n; i++ {
		path = append(path, i)
		backtrackCombine(n, k, path, ret, i+1)
		path = path[0 : len(path)-1]
	}
}
