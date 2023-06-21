package main

import "fmt"

func countNegatives(grid [][]int) int {
	i, j := len(grid)-1, len(grid[0])
	var offset int
	var count int
	for i >= 0 && offset < j {
		if grid[i][offset] < 0 {
			count += j - offset
			i--
		} else {
			end := j - 1
			for offset <= end {
				mid := (end-offset)/2 + offset
				if grid[i][mid] < 0 {
					end = mid - 1
				} else {
					offset = mid + 1
				}
			}
			count += j - offset
			i--
		}
	}

	return count
}

func main() {
	grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
	fmt.Println(countNegatives(grid))
}
