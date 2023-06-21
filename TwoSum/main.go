package main

func twoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)
	for i, n := range nums {
		diff := target - n
		if j, ok := prevMap[diff]; ok {
			return []int{i, j}
		}
		prevMap[n] = i
	}
	return nil
}
