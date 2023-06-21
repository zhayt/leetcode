package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	top, bottom, left, right := 0, n-1, 0, n-1
	elm := 1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			res[top][i] = elm
			elm++
		}
		top++
		for i := top; i <= bottom; i++ {
			res[i][right] = elm
			elm++
		}
		right--
		for i := right; i >= left; i-- {
			res[bottom][i] = elm
			elm++
		}
		bottom--
		for i := bottom; i >= top; i-- {
			res[i][left] = elm
			elm++
		}
		left++
	}
	return res
}
