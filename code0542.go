/*
给定一个由 0 和 1 组成的矩阵，找出每个元素到最近的 0 的距离。

两个相邻元素间的距离为 1 。

示例 1:
输入:

0 0 0
0 1 0
0 0 0
输出:

0 0 0
0 1 0
0 0 0
示例 2:
输入:

0 0 0
0 1 0
1 1 1
输出:

0 0 0
0 1 0
1 2 1
注意:

给定矩阵的元素个数不超过 10000。
给定矩阵中至少有一个元素是 0。
矩阵中的元素只在四个方向上相邻: 上、下、左、右。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/01-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

import "math"

func updateMatrix(matrix [][]int) [][]int {
	m := len(matrix)
	if m == 0 {
		return matrix
	}
	n := len(matrix[0])
	if n == 0 {
		return matrix
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != 0 {
				matrix[i][j] = math.MaxInt32
			}
		}
	}

	//visit := make([]bool, m*n)
	queue := NewLinkedQueue()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				queue.Add([3]int{0, i, j})
				//visit[i*n+j] = true
			}
		}
	}

	for queue.Size() != 0 {
		_ele, _ := queue.Peek()
		queue.Remove()
		ele := _ele.([3]int)
		v, x, y := ele[0], ele[1], ele[2]
		if x-1 >= 0 && matrix[x-1][y] != 0 && v+1 < matrix[x-1][y] {
			queue.Add([3]int{v + 1, x - 1, y})
			matrix[x-1][y] = v + 1
		}
		if x+1 < m && matrix[x+1][y] != 0 && v+1 < matrix[x+1][y] {
			queue.Add([3]int{v + 1, x + 1, y})
			matrix[x+1][y] = v + 1
		}
		if y-1 >= 0 && matrix[x][y-1] != 0 && v+1 < matrix[x][y-1] {
			queue.Add([3]int{v + 1, x, y - 1})
			matrix[x][y-1] = v + 1
		}
		if y+1 < n && matrix[x][y+1] != 0 && v+1 < matrix[x][y+1] {
			queue.Add([3]int{v + 1, x, y + 1})
			matrix[x][y+1] = v + 1
		}
	}

	return matrix
}
