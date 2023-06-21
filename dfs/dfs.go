package main

import "fmt"

var count int

func numEnclaves(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			flag := false
			if grid[i][j] == 1 {
				dfs(grid, i, j, &flag)
				if !flag {
					res += count
				}
				count = 0
			}
		}
	}
	return res
}

func dfs(grid [][]int, x, y int, flag *bool) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		*flag = true
		return
	}
	if grid[x][y] == 0 {
		return
	}
	grid[x][y] = 0
	count++
	dfs(grid, x-1, y, flag)
	dfs(grid, x+1, y, flag)
	dfs(grid, x, y-1, flag)
	dfs(grid, x, y+1, flag)
}

func main() {
	fmt.Println(numEnclaves([][]int{{
		0, 0, 0, 0,
	}, {
		1, 0, 1, 0,
	}, {
		0, 1, 1, 0,
	}, {0, 0, 0, 0}}))
}
