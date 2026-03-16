package main

import "fmt"

func isValid(s string) bool {
	strb := []byte(s)
	l := len(strb)
	stack := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		if strb[i] == '(' || strb[i] == '{' || strb[i] == '[' {
			stack = append(stack, strb[i])
		} else if strb[i] == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:(len(stack) - 1)]
			} else {
				stack = append(stack, strb[i])
			}

		} else if strb[i] == '}' {
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:(len(stack) - 1)]
			} else {
				stack = append(stack, strb[i])
			}

		} else if strb[i] == ']' {
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:(len(stack) - 1)]
			} else {
				stack = append(stack, strb[i])
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println("()", isValid("()"))
	fmt.Println("()[]{}", isValid("()[]{}"))
	fmt.Println("(]", isValid("(]"))
	fmt.Println("([])", isValid("([])"))
	fmt.Println("([)]", isValid("([)]"))
}
