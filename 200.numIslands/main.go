package main

import (
	"fmt"
)

func dfs(grid [][]byte, r, c int) {
	grid[r][c] = '0'
	m := len(grid)
	n := len(grid[0])
	if r-1 >= 0 && grid[r-1][c] == '1' {
		dfs(grid, r-1, c)
	}
	if r+1 < m && grid[r+1][c] == '1' {
		dfs(grid, r+1, c)
	}
	if c-1 >= 0 && grid[r][c-1] == '1' {
		dfs(grid, r, c-1)
	}
	if c+1 < n && grid[r][c+1] == '1' {
		dfs(grid, r, c+1)
	}
}

func numIslands(grid [][]byte) int {
	// 深度搜索 DFS
	// m := len(grid)
	// n := len(grid[0])
	// cnt := 0

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if grid[i][j] == '1' {
	// 			cnt++
	// 			dfs(grid, i, j)
	// 		}
	// 	}

	// }

	// return cnt

	// 广度搜索 bfs
	// m := len(grid)
	// n := len(grid[0])
	// cnt := 0
	// queue := list.New()

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if grid[i][j] == '1' {
	// 			cnt++
	// 			queue.PushBack(Postion{i, j})
	// 			for {
	// 				e := queue.Front()
	// 				if e == nil {
	// 					break
	// 				}
	// 				pos := e.Value.(Postion)
	// 				queue.Remove(e)

	// 				if grid[pos.r][pos.c] == '1' {
	// 					if pos.r-1 >= 0 && grid[pos.r-1][pos.c] == '1' {
	// 						queue.PushBack(Postion{pos.r - 1, pos.c})
	// 					}
	// 					if pos.r+1 < m && grid[pos.r+1][pos.c] == '1' {
	// 						queue.PushBack(Postion{pos.r + 1, pos.c})
	// 					}
	// 					if pos.c-1 >= 0 && grid[pos.r][pos.c-1] == '1' {
	// 						queue.PushBack(Postion{pos.r, pos.c - 1})
	// 					}
	// 					if pos.c+1 < n && grid[pos.r][pos.c+1] == '1' {
	// 						queue.PushBack(Postion{pos.r, pos.c + 1})
	// 					}
	// 					grid[pos.r][pos.c] = '0'
	// 				}

	// 			}

	// 		}

	// 	}
	// }

	// return cnt

	//并查集
	m := len(grid)
	n := len(grid[0])

	// 初始化
	parents := make([]int, 0, m*n)
	level := make([]int, 0, m*n)
	cnt := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				parents = append(parents, i*n+j)
				//fmt.Println(parents, i, j)
				cnt++
			} else {
				parents = append(parents, -1)
			}
			level = append(level, 1)
		}
	}

	// fmt.Println(parents)
	// fmt.Println(level)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				grid[i][j] = '0'
				if i-1 >= 0 && grid[i-1][j] == '1' {
					if unite(parents, level, i*n+j, (i-1)*n+j) {
						cnt--
					}
				}

				if i+1 < m && grid[i+1][j] == '1' {
					if unite(parents, level, i*n+j, (i+1)*n+j) {
						cnt--
					}
				}

				if j-1 >= 0 && grid[i][j-1] == '1' {
					if unite(parents, level, i*n+j, i*n+j-1) {
						cnt--
					}
				}

				if j+1 < n && grid[i][j+1] == '1' {
					if unite(parents, level, i*n+j, i*n+j+1) {
						cnt--
					}
				}
			}
		}
	}

	return cnt

}

type Postion struct {
	r int
	c int
}

func unite(parents, level []int, x, y int) bool {
	rootx := find(parents, x)
	rooty := find(parents, y)
	if rootx != rooty {
		if level[rootx] > level[rooty] {
			parents[rooty] = rootx
		} else if level[rootx] < level[rooty] {
			parents[rootx] = rooty
		} else {
			parents[rooty] = rootx
			level[rootx]++
		}
		return true
	}
	return false

}

func find(parents []int, i int) int {
	if parents[i] == i {
		return i
	}
	parents[i] = find(parents, parents[i])
	return parents[i]
}

func main() {
	m1 := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}}
	fmt.Println(numIslands(m1))

	m2 := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	fmt.Println(numIslands(m2))

}
