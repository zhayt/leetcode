package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	res := 0
	symbol := 1
	f := false
	for _, val := range s {
		if val == '-' && !f && res == 0 {
			f = true
			symbol = -1
			continue
		}
		if val == '+' && !f && res == 0 {
			f = true
			continue
		}
		if val == ' ' && res == 0 && !f {
			continue
		}
		if val >= '0' && val <= '9' {
			f = true
			res = res*10 + int(val-48)
			if res*symbol < math.MinInt32 {
				return math.MinInt32
			} else if res*symbol > math.MaxInt32 {
				return math.MaxInt32
			} else {
				continue
			}
		}
		if res == 0 {
			return 0
		} else {
			return symbol * res
		}
	}
	return res * symbol
}

func main() {
	fmt.Println(myAtoi("   +123"))
	fmt.Println(math.MaxInt32)
}
