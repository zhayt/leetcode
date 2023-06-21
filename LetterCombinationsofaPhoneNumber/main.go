package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	count := 0
	answer := ""
	for i >= 0 && j >= 0 {
		if a[i] == '1' && b[j] == '1' {
			if count != 0 {
				answer = "1" + answer
			} else {
				answer = "0" + answer
				count++
			}
		} else if a[i] != b[j] {
			if count != 0 {
				answer = "0" + answer
			} else {
				answer = "1" + answer
			}
		} else {
			if count != 0 {
				answer = "1" + answer
				count--
			} else {
				answer = "0" + answer
			}
		}
		i--
		j--
	}
	for ; i >= 0; i-- {
		if count != 0 && a[i] == '1' {
			answer = "0" + answer
			continue
		}
		if count != 0 && a[i] != '1' {
			answer = "1" + answer
			count--
			continue
		}
		answer = string(a[i]) + answer
	}
	for ; j >= 0; j-- {
		if count != 0 && b[j] == '1' {
			answer = "0" + answer
			continue
		}
		if count != 0 && b[j] != '1' {
			answer = "1" + answer
			count--
			continue
		}
		answer = string(b[j]) + answer
	}
	if count != 0 {
		answer = "1" + answer
	}
	return answer
}

func main() {
	fmt.Println(addBinary("110010", "10111"))
}
