package main

import (
	"fmt"
)

func partitionLabels(s string) []int {
	l := len(s)
	s_b := []byte(s)
	ret := []int{}
	char_map := make(map[byte]int, 26)
	for i := 0; i < l; i++ {
		char_map[s_b[i]] = i
	}

	start := 0
	max_postion := 0
	for i := 0; i < l; i++ {
		if char_map[s_b[i]] > max_postion {
			max_postion = char_map[s_b[i]]
		}

		if i == max_postion {
			ret = append(ret, max_postion-start+1)
			start = max_postion + 1
		}
	}
	return ret
}

func main() {
	ret := partitionLabels("ababcbacadefegdehijhklij")
	fmt.Println(ret)

	ret = partitionLabels("eccbbbbdec")
	fmt.Println(ret)

	ret = partitionLabels("qiejxqfnqceocmy")
	fmt.Println(ret)
}
