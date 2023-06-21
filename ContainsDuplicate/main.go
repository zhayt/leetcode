package main

import (
	"fmt"
	"sort"
)

// solution O(n log n)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// solution O(n)
func containsDuplicateUsingHash(nums []int) bool {
	set := make(map[int]bool)
	for _, n := range nums {
		if _, ok := set[n]; ok {
			return true
			break
		}
		set[n] = true
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 5, 7, 2, 3, 9, 4, 1}))
}
