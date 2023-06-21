package main

import "fmt"

func maxValue(n int, index int, maxSum int) int {
	res := maxSum / n
	maxSum = maxSum % n

	leftPtr := max(index-res, -1)
	rightPtr := min(index+res, n)

	if res != 1 {
		maxSum += (index - leftPtr) * (index - leftPtr - 1) / 2
		maxSum += (rightPtr - index) * (rightPtr - index - 1) / 2
		maxSum += (res - 1) * (leftPtr + 1)
		maxSum += (res - 1) * (n - rightPtr)
	}

	for maxSum > 0 && (leftPtr > -1 || rightPtr < n) {
		maxSum -= index - max(leftPtr, -1)
		maxSum -= min(rightPtr, n) - index - 1
		if maxSum >= 0 {
			res++
		}
		leftPtr--
		rightPtr++
	}

	res += maxSum / n

	return res
}

func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxValue(6, 1, 10))
}
