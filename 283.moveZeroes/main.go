package main

import "fmt"

func moveZeroes(nums []int) []int {
	zero_point := 0
	not_zero_point := 0
	once_flag := true
	for i := 0; i < len(nums); i++ {
		if once_flag {
			if nums[i] == 0 {
				zero_point = i
				once_flag = false
			}
		}
		if nums[i] != 0 {
			not_zero_point = i
			if !once_flag && zero_point < not_zero_point {
				nums[zero_point] = nums[not_zero_point]
				nums[not_zero_point] = 0
				zero_point++
			}
		}

	}
	return nums
}
func main() {
	nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	nums_ret := moveZeroes(nums)
	fmt.Println(nums_ret)

	nums1 := []int{1, 0}
	nums_ret1 := moveZeroes(nums1)
	fmt.Println(nums_ret1)
}
