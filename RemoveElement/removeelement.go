package main

import "fmt"

func removeElement(nums []int, val int) int {
	last := len(nums) - 1
	count := 0
	for i := 0; i < len(nums); i++ {
		if val == nums[i] {
			count++
			nums[i] = nums[last]
			nums[last] = val - 1
			i--
			last--
		}
	}
	return count
}

func main() {
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
