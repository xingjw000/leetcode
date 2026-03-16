package main

import (
	"fmt"
)

// 示例 1：

// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 示例 2：

// 输入：digits = "2"
// 输出：["a","b","c"]

func letterCombinations(digits string) []string {
	phone := map[int][]byte{}
	ans := []string{}

	for i := 2; i < 7; i++ {
		for j := 0; j < 3; j++ {
			phone[i] = append(phone[i], byte('a'+(i-2)*3+j))
		}
	}
	phone[7] = []byte("pqrs")
	phone[8] = []byte("tuv")
	phone[9] = []byte("wxzy")

	ret := []byte{}

	var dfs func(index int)

	dfs = func(index int) {
		if index == len(digits) {
			ans = append(ans, string(ret))
			return
		}

		num := int(digits[index] - '0')
		for i := 0; i < len(phone[num]); i++ {
			ret = append(ret, phone[num][i])
			dfs(index + 1)
			ret = ret[:len(ret)-1]
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("2"))
}
