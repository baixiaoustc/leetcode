/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func permute(nums []int) [][]int {
	path := make([]int, 0)
	ret := make([][]int, 0)
	visit := make([]bool, len(nums))
	backtrackPermute(nums, path, &ret, visit)
	return ret
}

func backtrackPermute(nums []int, path []int, ret *[][]int, visit []bool) {
	if len(path) == len(nums) {
		_path := make([]int, len(path))
		copy(_path, path)
		*ret = append(*ret, _path)
		return
	}

	for i, num := range nums {
		if visit[i] {
			continue
		}

		visit[i] = true
		path = append(path, num)
		backtrackPermute(nums, path, ret, visit)
		visit[i] = false
		path = path[:len(path)-1]
	}
}
