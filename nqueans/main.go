package main

import "fmt"

func solveNQueens(n int) [][]string {
	res := make([][]string, 0, n)
	board := make([][2]int, 0, n)

	placeQueen(0, n, board, &res)

	return res
}

func placeQueen(x, y int, board [][2]int, res *[][]string) {
	if x == y {
		addToResults(board, y, res)
		return
	}

	for row := 1; row <= y; row++ {
		if isValid(board, row, x+1) {
			// put quean
			board = append(board, [2]int{row, x + 1})
			placeQueen(x+1, y, board, res)
			board = board[:len(board)-1]
		}
	}
}

func isValid(board [][2]int, x, y int) bool {
	if len(board) == 0 {
		return true
	}

	for _, pos := range board {
		if pos[0] == x || pos[1] == y {
			return false
		}

		if x+y == pos[0]+pos[1] || x+pos[1] == y+pos[0] {
			return false
		}
	}

	return true
}

func addToResults(board [][2]int, y int, res *[][]string) {
	rows := []string{}

	for _, r := range board {
		row := buildRow(y)
		row[r[0]-1] = 'Q'
		rows = append(rows, string(row))
	}

	*res = append(*res, rows)
}

func buildRow(y int) []rune {
	row := make([]rune, y)

	for i := 0; i < y; i++ {
		row[i] = '.'
	}

	return row
}

func main() {
	fmt.Println(solveNQueens(4))
}
