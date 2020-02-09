/*
给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。



示例:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

输出:  [1,2,4,7,5,3,6,8,9]

解释:



说明:

给定矩阵中的元素总数不会超过 100000 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/diagonal-traverse
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func findDiagonalOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	} else if m == 1 {
		return matrix[0]
	}

	n := len(matrix[0])

	ret := make([]int, 0)
	//每一轮，横纵坐标之和一致，从0到m+n-1
	for i := 0; i < m+n-1; i++ {
		if i%2 == 0 {
			x, y := 0, 0
			if i < m {
				x = i
			} else {
				x = m - 1
			}
			y = i - x

			for x >= 0 && y < n {
				ret = append(ret, matrix[x][y])
				x--
				y++
			}
		} else {
			x, y := 0, 0
			if i < n {
				y = i
			} else {
				y = n - 1
			}
			x = i - y

			for y >= 0 && x < m {
				ret = append(ret, matrix[x][y])
				x++
				y--
			}
		}
	}

	return ret
}
