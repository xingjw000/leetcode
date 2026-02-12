package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	rets := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := nums[i] * (-1)
		k := n - 1

		for j := i + 1; j < n; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[j]+nums[k] > target {
				k--
			}

			if j == k {
				break
			}

			if nums[j]+nums[k] == target {
				rets = append(rets, []int{nums[i], nums[j], nums[k]})
			}
		}

	}

	return rets
}

func main() {
	test1 := []int{-1, 0, 1, 2, -1, -4}
	ret1 := threeSum(test1)

	fmt.Println(ret1)

	test2 := []int{0, 1, 1}
	ret2 := threeSum(test2)

	fmt.Println(ret2)

	test3 := []int{0, 0, 0}
	ret3 := threeSum(test3)

	fmt.Println(ret3)
}
