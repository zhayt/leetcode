package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	ret, l := 0, 0
	mark := [128]int{}
	for i := 0; i < len(s); i++ {
		index := mark[s[i]]
		if index > 0 {
			l = max(l, index)
		}

		mark[s[i]] = i + 1
		ret = max(ret, i+1-l)
	}
	return ret
}

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
