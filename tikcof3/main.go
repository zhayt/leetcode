package main

import "fmt"

func main() {
	var n, m, q int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Scan(&q)

	matrix := make([][]int, n)
	for i := 0; i < len(matrix); i++ {
		tmp := make([]int, m)
		matrix[i] = tmp
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var num int
			fmt.Scan(&num)
			matrix[i][j] = num
		}
	}

	res := make([]int, 0, q)

	for i := 0; i < q; i++ {
		var x, y, k int
		fmt.Scan(&x)
		fmt.Scan(&y)
		fmt.Scan(&k)

		copyMatrix := make([][]int, len(matrix))
		for i := 0; i < len(matrix); i++ {
			copyMatrix[i] = make([]int, len(matrix[i]))
			copy(copyMatrix[i], matrix[i])
		}
		var step int
		dfs(x-1, y-1, k, copyMatrix, &step)
		res = append(res, step--)
	}

	for _, v := range res {
		fmt.Println(v)
	}
}

func dfs(i, j, k int, matrix [][]int, step *int) {
	if k < 0 || matrix[i][j] == -1 {
		return
	}
	matrix[i][j] = -1
	*step++
	for l := 0; l < len(matrix); l++ {
		dfs(l, j, k-matrix[l][j], matrix, step)
	}

	for l := 0; l < len(matrix[0]); l++ {
		dfs(i, l, k-matrix[i][l], matrix, step)
	}
}
