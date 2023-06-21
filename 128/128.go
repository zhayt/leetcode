package main

func isValidSudoku(board [][]string) bool {
	for i := 0; i < 9; i++ {
		if i%3 == 0 {
			if checkSquare(board, i) {
				return false
			}
		}
	}

	return true
}

func checkColumn(board [][]string, i int) bool {
	hash := make(map[string]int)
	for j := 0; j < 9; j++ {
		if v, ok := hash[board[i][j]]; ok && v != '.' {
			return true
		}
		hash[board[i][j]]++
	}
	return false
}

func checkRow(board [][]string, j int) bool {
	hash := make(map[string]int)
	for i := 0; i < 9; j++ {
		if v, ok := hash[board[i][j]]; ok && v != '.' {
			return true
		}
		hash[board[i][j]]++
	}
	return false
}

func checkSquare(board [][]string, j int) bool {
	for i := 0; i < 9; i += 3 {
		hash := make(map[string]int)
		for ; j < j+2; j++ {
			for k := 0; k < 3; k++ {
				if _, ok := hash[board[j+i][k]]; ok && board[j+i][k] != "." {
					return true
				}
				hash[board[j+i][k]]++
			}
		}
	}
	return false
}

func main() {
	isValidSudoku([][]string{{"5", "3", ".", ".", "7", ".", ".", "1", "."},
		{"6", ".", ".", "1", "9", "5", ".", ".", "."},
		{".", "9", "8", ".", ".", ".", ".", ".", "1"},
		{"8", ".", ".", ".", "6", ".", ".", ".", "3"},
		{"4", ".", ".", "8", ".", "3", ".", ".", "1"},
		{"7", ".", ".", ".", "2", ".", ".", ".", "6"},
		{".", "6", ".", ".", ".", ".", "2", "8", "."}, {".", ".", ".", "4", "1", "9", ".", ".", "5"}, {".", ".", ".", ".", "8", ".", ".", "7", "9"}})
}
