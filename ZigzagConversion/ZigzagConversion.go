package main

import (
	"fmt"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) == 1 {
		return s
	}
	var result = make([]byte, len(s))
	index := 0
	step := numRows*2 - 2
	for i := 0; i < numRows && i < len(s); i++ {
		step2 := (numRows - i - 1) * 2
		for j := i; j < len(s); j += step {
			result[index] = s[j]
			index++
			if i == 0 || i == numRows-1 {
				continue
			}
			x := j + step2
			if x < len(s) {
				result[index] = s[x]
				index++
			}
		}
	}
	return string(result)
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
