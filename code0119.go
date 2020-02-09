/*
给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 3
输出: [1,3,3,1]
进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func getRow(rowIndex int) []int {
	tmp := make([]int, rowIndex+1)
	ret := make([]int, rowIndex+1)

	tmp[0] = 1
	ret[0] = 1
	for i := 1; i <= rowIndex; i++ {
		ret[i] = 1
		for j := 1; j < i; j++ {
			ret[j] = tmp[j-1] + tmp[j]
		}
		copy(tmp, ret)
	}

	return ret
}
