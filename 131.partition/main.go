package main

import (
	"fmt"
)

// 给你一个字符串 s，请你将 s 分割成一些 子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。

// 示例 1：

// 输入：s = "aab"
// 输出：[["a","a","b"],["aa","b"]]
// 示例 2：

// 输入：s = "a"
// 输出：[["a"]]

func partition(s string) (ans [][]string) {
	var dfs func(index int)

	dfs = func(index int) {
		if index == 0 {
			ans = append(ans, []string{string(s[index])})
			return
		}
		dfs(index - 1)
		m := len(ans)
		for i := 0; i < m; i++ {
			if len(ans[i]) == 1 {
				tmp := ans[i][len(ans[i])-1] + string(s[index])
				tmp_up := []string{}
				tmp_up = append(tmp_up, ans[i]...)
				if check(tmp) {
					tmp_up[len(ans[i])-1] = tmp
					ans = append(ans, tmp_up)
				}
			}

			if len(ans[i]) > 1 {
				tmp := ans[i][len(ans[i])-1] + string(s[index])
				tmp_up := []string{}
				tmp_up = append(tmp_up, ans[i]...)
				if check(tmp) {
					tmp_up[len(ans[i])-1] = tmp
					ans = append(ans, tmp_up)
				} else {
					tmp1 := ans[i][len(ans[i])-2] + tmp
					if check(tmp1) {
						tmp_up = tmp_up[:len(tmp_up)-1]
						tmp_up[len(tmp_up)-1] = tmp1
						ans = append(ans, tmp_up)
					}
				}

			}

			ans[i] = append(ans[i], string(s[index]))
		}

	}

	dfs(len(s) - 1)
	return ans

}

func check(str string) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
	fmt.Println(partition("aabaac"))
	fmt.Println(partition("a"))

	//fmt.Println(generateParenthesis(4))
}
