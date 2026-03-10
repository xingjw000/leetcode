package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	return partions(nums, 0, len(nums)-1, len(nums)-k)
}

func partions(nums []int, l, r, k int) int {
	fmt.Println(l, r, nums, nums[l])
	if l == r {
		return nums[l]
	}
	postion := l
	p_num := nums[postion]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < p_num; i++ {

		}

		for j--; nums[j] > p_num; j-- {

		}

		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}

	}
	fmt.Println(i, j, k, nums)
	if j < k {
		return partions(nums, j+1, r, k)
	} else {
		return partions(nums, l, j, k)
	}
}

func main() {
	test1 := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	ret1 := findKthLargest(test1, 4)

	fmt.Println(ret1)

	test2 := []int{3, 2, 1, 5, 6, 4}
	ret2 := findKthLargest(test2, 2)

	fmt.Println(ret2)

	// test3 := []int{-2, 1, 0, -2}
	// ret3 := subarraySum(test3, -1)

	// fmt.Println(ret3)
}
