package main

import "fmt"

func longestPalindrome(s string) string {
	answer := []int{0, 0}
	for i, _ := range s {
		start := i
		end := len(s) - 1
		if end-i < answer[1]-answer[0] {
			break
		}
		e := end
		n := i
		k := 0
		for start < end {
			if s[start] != s[end] {
				e--
				end = e
				start = i
				k = 0
			} else {
				start++
				if k == 0 {
					k = end
				}
				end--
			}
		}
		if answer[1]-answer[0] < k-n {
			answer[0] = n
			answer[1] = k
		}
	}
	return s[answer[0] : answer[1]+1]
}

func main() {
	fmt.Println(longestPalindrome("bacabab"))
}
