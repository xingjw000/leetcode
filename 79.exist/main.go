package main

import (
	"fmt"
)

// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCCED"
// 输出：true

// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "SEE"
// 输出：true

// 输入：board = [['A','B','C','E'],['S','F','C','S'],['A','D','E','E']], word = "ABCB"
// 输出：false

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	trace := make([]bool, m*n)
	var dfs func(index, i, j int) bool
	dfs = func(index, i, j int) bool {
		if word[index] != board[i][j] {
			return false
		}
		if index+1 == len(word) {
			return true
		}

		trace[i*n+j] = true

		if i+1 < m && !trace[(i+1)*n+j] {
			if dfs(index+1, i+1, j) {
				return true
			}
		}

		if i-1 >= 0 && !trace[(i-1)*n+j] {
			if dfs(index+1, i-1, j) {
				return true
			}
		}

		if j+1 < n && !trace[i*n+j+1] {
			if dfs(index+1, i, j+1) {
				return true
			}
		}

		if j-1 >= 0 && !trace[i*n+j-1] {
			if dfs(index+1, i, j-1) {
				return true
			}
		}
		trace[i*n+j] = false
		return false

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				if dfs(0, i, j) {
					return true
				}
			}

		}
	}

	return false
}

func main() {
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB"))
	fmt.Println(exist([][]byte{{'A', 'B', 'C', 'E'}}, "ABCE"))
	fmt.Println(exist([][]byte{{'a', 'a'}}, "aaa"))

	//fmt.Println(generateParenthesis(4))
}
