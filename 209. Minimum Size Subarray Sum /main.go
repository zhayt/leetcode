package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	var prevSum int
	res := len(nums)
	for i, j := 0, 0; (j < len(nums) || prevSum > target) && i < len(nums); {
		if prevSum < target {
			prevSum += nums[j]
			j++
		} else {
			res = min(res, j-i)
			prevSum -= nums[i]
			i++
		}
	}

	if res == len(nums) && prevSum < target {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(minSubArrayLen(15, arr))
}
