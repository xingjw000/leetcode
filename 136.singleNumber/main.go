package main

import "fmt"

func singleNumber(nums []int) int {
	// hash 算法
	// l := len(nums)
	// nums_map := make(map[int]bool, l/2+1)
	// for i := 0; i < l; i++ {
	// 	if !nums_map[nums[i]] {
	// 		nums_map[nums[i]] = true
	// 	} else {
	// 		delete(nums_map, nums[i])
	// 	}
	// }

	// for key, _ := range nums_map {
	// 	return key
	// }
	// return nums[0]

	// 异或
	ret := 0
	for _, num := range nums {
		ret ^= num
	}
	return ret
}

func main() {
	ret := singleNumber([]int{2, 2, 1})
	fmt.Println(ret)

	ret = singleNumber([]int{4, 1, 2, 1, 2})
	fmt.Println(ret)

	ret = singleNumber([]int{1})
	fmt.Println(ret)
}
