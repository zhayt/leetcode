package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	tmp := strings.ReplaceAll(strings.ToLower(s), " ", "")
	strings.EqualFold()
	i := 0
	j := len(tmp) - 1
	for i < j {
		if tmp[i] < '0' || tmp[i] > '9' && tmp[i] < 'a' || tmp[i] > 'z' {
			i++
			continue
		}
		if tmp[j] < '0' || tmp[j] > '9' && tmp[j] < 'a' || tmp[j] > 'z' {
			j--
			continue
		}
		if tmp[i] != tmp[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	str := "race a car"
	fmt.Println(isPalindrome(str))
}
