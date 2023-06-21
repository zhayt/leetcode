package main

import "fmt"

func main() {
	grid := [][]int{{2, 1, 3, 5}, {1, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}
	fmt.Println(maxMoves(grid))
}

// 4 + 3 / 6
// 1 2 3 4 5 6

func maxMoves(grid [][]int) int {
	return dfs(grid, grid[0][0])
}

func dfs(grid [][]bool, item int) int {

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
