package main

import (
	"fmt"
)

// 按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

// 输入：n = 4
// 输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
// 解释：如上图所示，4 皇后问题存在两个不同的解法。
// 示例 2：

// 输入：n = 1
// 输出：[["Q"]]

func solveNQueens(n int) (ans [][]string) {
	colums := make([]bool, n)
	diagonal_r := make([]bool, 2*n)
	diagonal_l := make([]bool, 2*n)
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}

	var backTrace func(index int)

	backTrace = func(index int) {
		if index == n {
			fmt.Println(queens)
			tmp := []string{}
			for i := 0; i < n; i++ {
				row := make([]byte, n)
				for j := 0; j < n; j++ {
					row[j] = '.'
				}
				row[queens[i]] = 'Q'
				tmp = append(tmp, string(row))
			}
			ans = append(ans, tmp)
			return
		}

		for i := 0; i < n; i++ {

			if colums[i] {
				continue
			}

			if diagonal_r[index-i+n] {
				continue
			}

			if diagonal_l[i+index] {
				continue
			}
			queens[index] = i
			colums[i] = true
			diagonal_r[index-i+n] = true
			diagonal_l[i+index] = true

			backTrace(index + 1)

			queens[index] = -1
			diagonal_r[index-i+n] = false
			diagonal_l[i+index] = false
			colums[i] = false
		}

	}

	backTrace(0)
	return ans
}

func main() {
	fmt.Println(solveNQueens(4))
	fmt.Println(solveNQueens(1))
	//fmt.Println(generateParenthesis(4))
}
