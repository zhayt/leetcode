package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 7, 1, 7, 5}
	nums2 := []int{1, 9, 2, 5, 1}
	fmt.Println(maxUncrossedLines(nums1, nums2))
}

func maxUncrossedLines(nums1 []int, nums2 []int) int {
	n1 := len(nums1)
	n2 := len(nums2)

	memo := make([][]int, n1+1)
	for i := range memo {
		memo[i] = make([]int, n2+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	return solve(n1, n2, nums1, nums2, memo)
}

func solve(i, j int, nums1, nums2 []int, memo [][]int) int {
	if i <= 0 || j <= 0 {
		return 0
	}

	if memo[i][j] != -1 {
		return memo[i][j]
	}

	if nums1[i-1] == nums2[j-1] {
		memo[i][j] = 1 + solve(i-1, j-1, nums1, nums2, memo)
	} else {
		memo[i][j] = max(solve(i, j-1, nums1, nums2, memo), solve(i-1, j, nums1, nums2, memo))
	}

	return memo[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
