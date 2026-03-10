package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {

	wordDictLen := len(wordDict)
	sLen := len(s)
	rets := make([]bool, sLen+1)
	rets[0] = true
	for i := 0; i <= sLen; i++ {
		for j := 0; j < wordDictLen; j++ {
			if i-len(wordDict[j]) >= 0 {
				if rets[i-len(wordDict[j])] {
					if s[i-len(wordDict[j]):i] == wordDict[j] {
						rets[i] = true
						break
					} else {
						rets[i] = false
					}
				}
			}
		}
	}
	return rets[sLen]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))

}
