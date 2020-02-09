/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

//此题的解法很多，值得思考
func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}

	n := len(matrix[0])
	ret := make([]int, 0)
	for m > 0 && n > 0 {
		ret = append(ret, matrix[0]...)
		m--
		if m <= 0 {
			break
		}

		tmp := make([][]int, m)
		for i := 0; i < m; i++ {
			tmp[i] = matrix[i+1]
		}

		//tmp is [m,n], trans is [n, m]
		trans := transpose(tmp)
		revers := make([][]int, n)
		for i := 0; i < n; i++ {
			revers[i] = trans[n-1-i]
		}

		matrix = revers
		m, n = n, m
	}

	return ret
}
