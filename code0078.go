/*
给定一组不含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: nums = [1,2,3]
输出:
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func subsets(nums []int) [][]int {
	path := make([]int, 0)
	ret := make([][]int, 0)
	backtrackSubsets(nums, path, &ret, 0)
	return ret
}

func backtrackSubsets(nums []int, path []int, ret *[][]int, start int) {
	_path := make([]int, len(path))
	copy(_path, path)
	*ret = append(*ret, _path)

	for i := start; i < len(nums); i++ {
		path = append(path, nums[i])
		backtrackSubsets(nums, path, ret, i+1)
		path = path[0 : len(path)-1]
	}
}
