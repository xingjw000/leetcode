package main

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
	rets := make([]int, n+1)
	for i := 1; i <= n; i++ {
		rets[i] = math.MaxInt
		for j := 1; j*j <= i; j++ {
			if j*j == i {
				rets[i] = 1
			} else {
				rets[i] = min(rets[i-j*j]+1, rets[i])

			}
		}
	}
	return rets[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))

}
