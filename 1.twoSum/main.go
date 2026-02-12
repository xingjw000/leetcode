package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		value, ok := numMap[target-num]

		if ok {
			return []int{value, i}
		}
		numMap[num] = i
	}
	return nil

}

func main() {
	test1 := []int{2, 7, 11, 15}
	target1 := 9
	ret := twoSum(test1, target1)
	fmt.Println(ret)

	test2 := []int{4, 5, 8, 12, 3, 3}
	target2 := 6
	ret = twoSum(test2, target2)
	fmt.Println(ret)
}
