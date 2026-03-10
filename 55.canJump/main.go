package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	// 后遍历
	// l := len(nums)
	// if l == 1 {
	// 	return true
	// }
	// need_min_steps := make([]int, l)
	// for i := 0; i < l; i++ {
	// 	need_min_steps[i] = 1
	// }
	// for i := l - 2; i > 0; i-- {
	// 	if nums[i] == 0 || nums[i] < need_min_steps[i] {
	// 		need_min_steps[i-1] += need_min_steps[i]
	// 	}
	// }

	// ret := false
	// //fmt.Println(need_min_steps)
	// for i := 0; i < l-1; i++ {
	// 	if nums[i] < need_min_steps[i] && ret == false {
	// 		return false
	// 	}

	// 	if nums[i] >= need_min_steps[i] {
	// 		ret = true
	// 	}
	// }
	// return ret

	max_reach := 0
	l := len(nums)

	for i, num := range nums {
		if i <= max_reach {
			if max_reach < num+i {
				max_reach = num + i
			}

			if max_reach >= l-1 {
				return true
			}
		} else {
			return false
		}
	}
	return false
}

func main() {
	ret := canJump([]int{2, 3, 1, 1, 4})
	fmt.Println(ret)

	ret = canJump([]int{3, 2, 1, 0, 4})
	fmt.Println(ret)

	ret = canJump([]int{2, 0, 0})
	fmt.Println(ret)
	ret = canJump([]int{0, 1})
	fmt.Println(ret)
	ret = canJump([]int{0, 2, 3})
	fmt.Println(ret)

	ret = canJump([]int{3, 0, 8, 2, 0, 0, 1})
	fmt.Println(ret)
}
