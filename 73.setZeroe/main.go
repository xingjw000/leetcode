package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	// m := len(matrix)
	// n := len(matrix[0])

	// zero_m := make([]int, 0, m)
	// zero_n := make([]int, 0, n)

	// for i := 0; i < m; i++ {
	// 	for j := 0; j < n; j++ {
	// 		if matrix[i][j] == 0 {
	// 			zero_m = append(zero_m, i)
	// 			zero_n = append(zero_n, j)
	// 		}
	// 	}
	// }

	// for i := 0; i < len(zero_m); i++ {
	// 	for j := 0; j < n; j++ {
	// 		matrix[zero_m[i]][j] = 0
	// 	}
	// }

	// for i := 0; i < len(zero_n); i++ {
	// 	for j := 0; j < m; j++ {
	// 		matrix[j][zero_n[i]] = 0
	// 	}
	// }

	m := len(matrix)
	n := len(matrix[0])
	zero_flag := false

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			zero_flag = true
		}

		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := m - 1; i >= 0; i-- {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}

		if zero_flag {
			matrix[i][0] = 0
		}
	}

}

func main() {
	m1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(m1)
	fmt.Println(m1)

	m2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(m2)
	fmt.Println(m2)

}
