package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	ret := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		pre := ret[i-1]
		curr := make([]int, i+1)
		curr[0] = 1
		for j := 1; j < i; j++ {
			curr[j] = pre[j-1] + pre[j]
		}
		curr[i] = 1
		ret = append(ret, curr)
	}
	return ret
}

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))

}
