package main

import "fmt"

func longestConsecutive(nums []int) int {
	map_nums := make(map[int]bool)
	for _, v := range nums {
		map_nums[v] = true
	}

	longest_num := 0

	for v := range map_nums {
		if !map_nums[v-1] {
			curr_longest_num := 1
			curr_num := v
			for map_nums[curr_num+1] {
				curr_longest_num++
				curr_num++
			}
			if longest_num < curr_longest_num {
				longest_num = curr_longest_num
			}
		}
	}
	return longest_num

	// numSet := map[int]bool{}
	// for _, num := range nums {
	// 	numSet[num] = true
	// }
	// longestStreak := 0
	// for num := range numSet {
	// 	if !numSet[num-1] {
	// 		currentNum := num
	// 		currentStreak := 1
	// 		for numSet[currentNum+1] {
	// 			currentNum++
	// 			currentStreak++
	// 		}
	// 		if longestStreak < currentStreak {
	// 			longestStreak = currentStreak
	// 		}
	// 	}
	// }
	// return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	longest_num := longestConsecutive(nums)
	fmt.Println(longest_num)
}
