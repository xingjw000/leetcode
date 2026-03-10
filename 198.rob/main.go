package main

import (
	"fmt"
)

func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	if l == 2 {
		return max(nums[0], nums[1])
	}

	prePre := nums[0]
	pre := max(nums[0], nums[1])
	for i := 2; i < l; i++ {
		curr := max(prePre+nums[i], pre)
		prePre = pre
		pre = curr
	}
	return pre
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1, 14, 15, 2}))
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
