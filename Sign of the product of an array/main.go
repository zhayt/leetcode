package main

import (
	"fmt"
	"sort"
)

func arraySign(nums []int) int {
	sort.Ints(nums)

	left := 0
	right := len(nums)
	mid := 0
	for left <= right {
		mid = (right-left)/2 + left
		if nums[mid] == 0 {
			return 0
		} else if nums[mid] > 0 {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if right-1%2 == 0 {
		return 1
	}

	return -1
}

func main() {
	arr := []int{41, 65, 14, 80, 20, 10, 55, 58, 24, 56, 28, 86, 96, 10, 3, 84, 4, 41, 13, 32, 42, 43, 83, 78, 82, 70, 15, -41}
	fmt.Println(arraySign(arr))
}
