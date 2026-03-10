package main

import (
	"fmt"
)

func jump(nums []int) int {

	max_reach := 0
	l := len(nums)
	cnt := 0
	end := 0

	for i := 0; i < l-1; i++ {
		num := nums[i]
		if i+num > max_reach {
			max_reach = i + num
		}

		if i == end {
			end = max_reach
			cnt++
		}
	}
	return cnt
}

func main() {
	ret := jump([]int{2, 3, 1, 1, 4})
	fmt.Println(ret)

	ret = jump([]int{2, 3, 0, 1, 4})
	fmt.Println(ret)
}
