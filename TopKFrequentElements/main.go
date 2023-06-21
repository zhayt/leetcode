package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	set := make(map[int]int)
	for _, n := range nums {
		set[n]++
	}
	s := make([][]int, len(nums)+1)

	for key, value := range set {
		s[value] = append(s[value], key)
	}

	answer := make([]int, 0, k)
	i := len(s) - 1
	for k != 0 && i > 0 {
		if s[i] != nil {
			answer = append(answer, s[i]...)
			k = k - len(s[i])
		}
		i--
	}
	return answer
}

func main() {
	fmt.Println(topKFrequent([]int{1, 2}, 3))
}
