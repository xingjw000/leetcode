package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	l := len(nums)
	cnt := 0
	pre_sum_map := make(map[int]int)
	pre_sum_map[0] = 1
	pre_sum := 0
	for i := 0; i < l; i++ {
		pre_sum += nums[i]
		if _, ok := pre_sum_map[pre_sum-k]; ok {
			cnt += pre_sum_map[pre_sum-k]
		}
		pre_sum_map[pre_sum]++
	}
	return cnt
}

func main() {
	test1 := []int{1, 1, 1}
	ret1 := subarraySum(test1, 2)

	fmt.Println(ret1)

	test2 := []int{1, 2, 3}
	ret2 := subarraySum(test2, 3)

	fmt.Println(ret2)

	test3 := []int{-2, 1, 0, -2}
	ret3 := subarraySum(test3, -1)

	fmt.Println(ret3)
}
