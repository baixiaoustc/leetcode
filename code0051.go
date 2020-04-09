/*
n 皇后问题研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。



上图为 8 皇后问题的一种解法。

给定一个整数 n，返回所有不同的 n 皇后问题的解决方案。

每一种解法包含一个明确的 n 皇后问题的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

示例:

输入: 4
输出: [
 [".Q..",  // 解法 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // 解法 2
  "Q...",
  "...Q",
  ".Q.."]
]
解释: 4 皇后问题存在两个不同的解法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-queens
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func solveNQueens(n int) [][]string {
	ret := make([][]string, 0)
	board := make([][]int, n) //board其实就是path
	for i := range board {
		board[i] = make([]int, n)
	}
	backtrackNQueen(board, &ret, 0)
	return ret
}

func backtrackNQueen(board [][]int, ret *[][]string, row int) {
	if row == len(board) {
		path := make([]string, len(board))
		for i := range board {
			var p string
			for _, n := range board[i] {
				if n == 0 {
					p += "."
				} else if n == 1 {
					p += "Q"
				}
			}
			path[i] = p
		}
		*ret = append(*ret, path)
		return
	}

	for col := 0; col < len(board); col++ {
		if !isValidQueue(board, row, col) {
			continue
		}
		board[row][col] = 1
		backtrackNQueen(board, ret, row+1)
		board[row][col] = 0
	}
}

func isValidQueue(board [][]int, row, col int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 1 && (i == row || j == col) {
				return false
			}
		}
	}

	//左上右上左下右下四个方向都不可用
	//实际只用检测上半部分
	for i, j := row-1, col+1; i >= 0 && j < len(board); {
		if board[i][j] == 1 {
			return false
		}
		i--
		j++
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; {
		if board[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	return true
}
