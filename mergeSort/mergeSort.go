package main

import "fmt"

func maxArea(height []int) int {
	var items = make(map[int][]int)
	for i, val := range height {
		items[val] = append(items[val], i)
	}
	result := 0
	var index []int
	sorted := mergeSort(height)
	for i := 0; i < len(sorted); i++ {
		if i != 0 {
			if sorted[i] == sorted[i-1] {
				continue
			}
		}
		for j := 0; j < len(items[sorted[i]]); j++ {
			index = append(index, items[sorted[i]][j])
		}
	}
	for i, val := range sorted {
		for j := i + 1; j < len(index); j++ {
			ras := index[i] - index[j]
			if ras < 0 {
				ras *= -1
			}
			if result < ras*val {
				result = ras * val
			}
		}
	}
	return result
}

func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []int, b []int) []int {
	var merged []int
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		merged = append(merged, a[i])
	}
	for ; j < len(b); j++ {
		merged = append(merged, b[j])
	}
	return merged
}

func main() {
	var arr = []int{1, 0, 0, 0, 0, 0, 0, 2, 2}
	fmt.Println(maxArea(arr))

}
