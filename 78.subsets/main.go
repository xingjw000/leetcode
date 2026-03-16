package main

import (
	"fmt"
)

// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// 示例 2：

// 输入：nums = [0]
// 输出：[[],[0]]

func subsets(nums []int) (ans [][]int) {
	// l := len(nums)
	// result := [][]int{}
	// result = append(result, []int{})

	// for i := 0; i < l; i++ {
	// 	backTrace(nums, &result, []int{}, i)
	// }
	// return result

	set := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		fmt.Println("pre", cur, set)
		dfs(cur + 1)
		set = set[:len(set)-1]
		fmt.Println("after", cur, set)
		dfs(cur + 1)
	}
	dfs(0)
	return

}

func backTrace(nums []int, result *[][]int, pre_ret []int, index int) []int {
	pre_ret = append(pre_ret, nums[index])
	*result = append(*result, append([]int{}, pre_ret...))

	for i := index + 1; i < len(nums); i++ {
		pre_ret = backTrace(nums, result, pre_ret, i)
		pre_ret = pre_ret[:len(pre_ret)-1]
	}

	return pre_ret
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{0}))
}
