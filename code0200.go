/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

//BFS
func numIslands(grid [][]byte) int {
	directions := [][2]int{[2]int{-1, 0}, [2]int{1, 0}, [2]int{0, -1}, [2]int{0, 1}}

	var m, n int
	if m = len(grid); m == 0 {
		return 0
	} else if n = len(grid[0]); n == 0 {
		return 0
	}

	queue := NewLinkedQueue()
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				queue.Add([2]int{i, j})
				grid[i][j] = 0

				for queue.Size() != 0 {
					ele, err := queue.Peek()
					if err != nil {
						panic("queue peek")
					}
					err = queue.Remove()
					if err != nil {
						panic("queue remove")
					}

					coordinate := ele.([2]int)
					for _, dir := range directions {
						ii := coordinate[0] + dir[0]
						jj := coordinate[1] + dir[1]
						if ii < 0 || ii >= m {
							continue
						}
						if jj < 0 || jj >= n {
							continue
						}
						if grid[ii][jj] == '1' {
							queue.Add([2]int{ii, jj})
							grid[ii][jj] = 0
						}
					}
				}
			}
		}
	}

	return count
}

//DFS
func numIslandsDFS(grid [][]byte) int {
	var m, n int
	if m = len(grid); m == 0 {
		return 0
	} else if n = len(grid[0]); n == 0 {
		return 0
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				islandsDFS(grid, i, j, m, n)
			}
		}
	}

	return count
}

func islandsDFS(grid [][]byte, i, j, m, n int) {
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	islandsDFS(grid, i-1, j, m, n)
	islandsDFS(grid, i+1, j, m, n)
	islandsDFS(grid, i, j-1, m, n)
	islandsDFS(grid, i, j+1, m, n)
}

//并查集
//没理解透彻
func numIslandsUnion(grid [][]byte) int {
	var m, n int
	if m = len(grid); m == 0 {
		return 0
	} else if n = len(grid[0]); n == 0 {
		return 0
	}

	uf := new(UnionFind)
	uf.Init(grid)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					uf.Union(i*n+j, (i-1)*n+j)
				}
				if i+1 < m && grid[i+1][j] == '1' {
					uf.Union(i*n+j, (i+1)*n+j)
				}
				if j-1 >= 0 && grid[i][j-1] == '1' {
					uf.Union(i*n+j, i*n+j-1)
				}
				if j+1 < n && grid[i][j+1] == '1' {
					uf.Union(i*n+j, i*n+j+1)
				}
			}
		}
	}

	return uf.GetCount()
}

type UnionFind struct {
	count  int
	parent []int
	rank   []int
}

func (uf *UnionFind) Init(grid [][]byte) {
	var m, n int
	if m = len(grid); m == 0 {
		return
	} else if n = len(grid[0]); n == 0 {
		return
	}

	uf.parent = make([]int, m*n)
	uf.rank = make([]int, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				uf.count++
				uf.parent[i*n+j] = i*n + j
			}
		}
	}
}

func (uf *UnionFind) Union(x, y int) {
	rootx := uf.Parent(x)
	rooty := uf.Parent(y)
	if rootx != rooty {
		if uf.rank[rootx] > uf.rank[rooty] {
			uf.parent[rooty] = rootx
		} else if uf.rank[rootx] < uf.rank[rooty] {
			uf.parent[rootx] = rooty
		} else {
			uf.parent[rooty] = rootx
			uf.rank[rootx] += 1
		}
		uf.count--
	}
}

func (uf *UnionFind) Parent(x int) int {
	if p := uf.parent[x]; p != x {
		uf.parent[x] = uf.Parent(p)
	}
	return uf.parent[x]
}

func (uf *UnionFind) GetCount() int {
	return uf.count
}
