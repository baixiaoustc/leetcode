/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func generate(numRows int) [][]int {
	ret := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		if i == 0 {
			row[0] = 1
		} else {
			row[0] = 1
			row[i] = 1
			for j := 1; j < i; j++ {
				row[j] = ret[i-1][j-1] + ret[i-1][j]
			}
		}

		ret[i] = row
	}

	return ret
}
