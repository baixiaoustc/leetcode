/*
给定一个可能包含重复元素的整数数组 nums，返回该数组所有可能的子集（幂集）。

说明：解集不能包含重复的子集。

示例:

输入: [1,2,2]
输出:
[
  [2],
  [1],
  [1,2,2],
  [2,2],
  [1,2],
  []
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/subsets-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	path := make([]int, 0)
	ret := make([][]int, 0)
	sort.Ints(nums)
	fmt.Println(nums)
	backtrackSubsetsWithDup(nums, path, &ret, 0)
	return ret
}

func backtrackSubsetsWithDup(nums []int, path []int, ret *[][]int, start int) {
	_path := make([]int, len(path))
	copy(_path, path)
	sort.Ints(_path)

	fmt.Println("in", _path, *ret)
	*ret = append(*ret, _path)

	for i := start; i < len(nums); i++ {
		//排序之后，以此去重
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		path = append(path, nums[i])
		backtrackSubsetsWithDup(nums, path, ret, i+1)
		path = path[0 : len(path)-1]
	}
}

//func subsetDuplicate(_path []int, ret *[][]int) bool {
//	fmt.Println("in", _path, *ret)
//
//	for _, r := range *ret {
//		fmt.Println(_path, r)
//		if len(r) != len(_path) {
//			continue
//		}
//
//		skip := false
//		for i := 0; i < len(_path); i++ {
//			if _path[i] != r[i] {
//				skip = true
//				break
//			}
//		}
//		if skip {
//			continue
//		}
//
//		fmt.Println("true")
//		return true
//	}
//
//	return false
//}
