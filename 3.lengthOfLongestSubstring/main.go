package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	s_bs := []byte(s)
	s_bs_len := len(s_bs)
	no_same_set := make(map[byte]bool, 0)
	left := 0
	right := 0
	longestsubstringlen := 0
	setlen := 0
	for right < s_bs_len {
		if !no_same_set[s_bs[right]] {
			no_same_set[s_bs[right]] = true
			right++
			setlen++

		} else {
			//delete(no_same_set, s_bs[left])
			no_same_set[s_bs[left]] = false
			setlen--
			left++
		}
		if longestsubstringlen < setlen {
			longestsubstringlen = setlen
		}
	}
	return longestsubstringlen
}

func main() {

	ret := lengthOfLongestSubstring("abcabcbb")
	fmt.Println(ret)

	ret = lengthOfLongestSubstring("bbbbb")
	fmt.Println(ret)

	ret = lengthOfLongestSubstring(" ")
	fmt.Println(ret)
}
