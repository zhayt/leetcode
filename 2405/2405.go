package main

import "fmt"

func partitionString(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		var letter [26]bool
		for j := i; j < len(s); j++ {
			if letter[s[j]-'a'] {
				i = j - 1
				res++
				break
			}
			letter[s[j]-'a'] = true
		}
	}

	return res + 1
}

func main() {
	fmt.Println(partitionString("abacaba"))
}
