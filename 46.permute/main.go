package main

import "fmt"

func permute(nums []int) [][]int {
	// result := [][]int{}
	// l := len(nums)

	// if l == 1 {
	// 	result = append(result, nums)
	// 	return result
	// }
	// result = append(result, []int{nums[0]})

	// for i := 1; i < l; i++ {
	// 	r_l := len(result)
	// 	for k := 0; k < r_l; k++ {
	// 		v := result[k]
	// 		v_l := len(v)
	// 		for j := 0; j < v_l; j++ {
	// 			v_new := make([]int, 0, v_l+1)
	// 			v_new = append(v_new, v[:j]...)
	// 			v_new = append(v_new, nums[i])
	// 			v_new = append(v_new, v[j:]...)
	// 			result = append(result, v_new)
	// 		}
	// 		result[k] = append(result[k], nums[i])
	// 	}
	// }
	// return result

	used := make([]bool, len(nums))
	result := [][]int{}
	out := make([]int, 0)
	backTrace(nums, used, out, &result)
	return result

}

func backTrace(nums []int, used []bool, out []int, result *[][]int) {
	if len(out) == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, out)
		*result = append(*result, tmp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if !used[i] {
			used[i] = true
			out = append(out, nums[i])
			backTrace(nums, used, out, result)
			out = out[:len(out)-1]
			used[i] = false
		}
	}

}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
