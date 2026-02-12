package main

import "fmt"

func searchInsert(nums []int, target int) int {
	// 递归
	// l := len(nums)
	// if l == 0 {
	// 	return 0
	// }
	// begin := 0
	// end := l - 1
	// postion := search2(begin, end, target, &nums)
	// return postion

	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left

}

func search2(begin, end, target int, nums *[]int) int {
	mid := (begin + end) / 2
	if (*nums)[mid] > target {
		end = mid
	} else if (*nums)[mid] < target {
		begin = mid
	} else {
		return mid
	}
	if begin+1 == end || begin == end {
		if target <= (*nums)[begin] {
			return begin
		}
		if target > (*nums)[begin] && target <= (*nums)[end] {
			return end
		}
		if target > (*nums)[end] {
			return end + 1
		}
	}

	return search2(begin, end, target, nums)
}

func main() {
	test1 := []int{1, 3, 5, 6}
	target1 := 5
	ret := searchInsert(test1, target1)
	fmt.Println(ret)

	test2 := []int{1, 3, 5, 6}
	target2 := 2
	ret = searchInsert(test2, target2)
	fmt.Println(ret)

	ret = searchInsert(test2, 7)
	fmt.Println(ret)

	ret = searchInsert([]int{1}, 0)
	fmt.Println(ret)
}
