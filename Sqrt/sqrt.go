package main

import "fmt"

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	for left <= right {
		mid := left + (right-left)/2
		if mid > x/mid {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}

func main() {
	fmt.Println(mySqrt(16))
}
