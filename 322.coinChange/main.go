package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	rets := make([]int, amount+2)
	rets[amount+1] = amount + 1
	for i := 1; i <= amount; i++ {
		minCnt := amount + 1
		coins_cnt := len(coins)
		for j := 0; j < coins_cnt; j++ {
			pre := i - coins[j]
			if pre < 0 {
				pre = amount + 1
			}

			if rets[pre] == -1 {
				continue
			}

			minCnt = min(rets[pre], minCnt)
		}
		if minCnt == amount+1 {
			rets[i] = -1
		} else {
			rets[i] = minCnt + 1
		}

	}
	return rets[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))

}
