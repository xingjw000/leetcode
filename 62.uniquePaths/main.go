package main

import (
	"fmt"
)

func uniquePaths(m int, n int) int {
	// dfs 超时了
	// ret := 0
	// step(m, n, 0, 0, &ret)
	// return ret

	data := make([]int, m*n)
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = data[i*n : (i+1)*n]
	}
	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if i == 0 && j == 0 {
				continue
			}
			var left, right int
			if i-1 >= 0 {
				left = dp[i-1][j]
			}

			if j-1 >= 0 {
				right = dp[i][j-1]
			}
			dp[i][j] = left + right
		}
	}

	return dp[m-1][n-1]
}

func step(m, n, x, y int, cnt *int) {
	if x == m-1 && y == n-1 {
		*cnt++
		return
	}
	if x < m-1 {
		step(m, n, x+1, y, cnt)
	}

	if y < n-1 {
		step(m, n, x, y+1, cnt)
	}
}

type Postion struct {
	x, y int
}

func main() {
	fmt.Println(uniquePaths(3, 7))

	fmt.Println(uniquePaths(3, 2))

}
