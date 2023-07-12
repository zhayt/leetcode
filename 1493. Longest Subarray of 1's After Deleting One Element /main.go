package main

import "fmt"

func main() {
	fmt.Println(longestSubarray([]int{1, 1, 1, 0, 1}))
}

func longestSubarray(nums []int) int {
	var prev, res, curr int
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 && curr != 0 {
			break
		}
		if nums[i] == 1 {
			curr++
		}
	}

	res = curr
	prev = curr
	curr = 0
	for ; i < len(nums); i++ {
		if nums[i] == 0 && curr != 0 {
			res = max(prev+curr, res)
			prev = curr
			curr = 0
			continue
		}

		if nums[i] == 1 {
			curr++
		}
	}

	return max(prev+curr, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
