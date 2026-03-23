package main

import "fmt"

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 示例 1：

// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：

// 输入：n = 1
// 输出：["()"]

func generateParenthesis(n int) (ans []string) {

	tmp_list := make([]byte, 2*n)
	tmp_list[0] = '('

	var bfs func(index int)
	bfs = func(index int) {
		if index == 2*n {
			if check(tmp_list) {
				ans = append(ans, string(tmp_list))
			}

			return
		}

		tmp_list[index] = '('
		bfs(index + 1)
		tmp_list[index] = ')'
		bfs(index + 1)
	}
	bfs(1)
	return ans
}

func check(str []byte) bool {
	stack := []byte{}
	for _, v := range str {
		if len(stack) == 0 || v == '(' {
			stack = append(stack, v)
		} else {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 && stack[0] == ')' {
			return false
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(1))
	//fmt.Println(generateParenthesis(4))
}
