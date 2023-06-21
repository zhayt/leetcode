package main

import "fmt"

func trap(height []int) int {
	i := 1
	j := len(height) - 1
	maxi := 0
	curn := height[0]
	sum := 0
	msum := 0
	for i <= j {
		if curn <= height[i] {
			maxi = i
			sum += msum
			msum = 0
			curn = height[i]
			i++
			continue
		}
		msum = msum + (curn - height[i])
		i++
	}
	msum = 0
	curn = height[j]
	for maxi <= j-1 {
		if curn <= height[j] {
			sum += msum
			msum = 0
			curn = height[j]
			j--
			continue
		}
		msum = msum + (curn - height[j])
		j--
	}
	sum += msum
	return sum
}

func main() {
	numbers := []int{4, 2, 3}
	fmt.Println(trap(numbers))
}
