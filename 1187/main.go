package main

import (
	"fmt"
	"math"
	"sort"
)

func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	dp := make(map[Pair]int)
	var dfs func(i, prev int) int
	dfs = func(i, prev int) int {
		if i == len(arr1) {
			return 0
		}
		p := Pair{i, prev}
		if v, ok := dp[p]; ok {
			return v
		}
		cost := math.MaxInt32
		if arr1[i] > prev {
			cost = dfs(i+1, arr1[i])
		}
		idx := sort.Search(len(arr2), func(i int) bool { return arr2[i] > prev })
		if idx < len(arr2) {
			cost = min(cost, 1+dfs(i+1, arr2[idx]))
		}
		dp[p] = cost
		return cost
	}
	ans := dfs(0, -1)
	if ans == math.MaxInt32 {
		return -1
	}
	return ans
}

type Pair struct {
	a, b int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	arr1 := []int{1, 5, 3, 6, 7}
	arr2 := []int{1, 3, 2, 4}
	fmt.Println(makeArrayIncreasing(arr1, arr2))
}
