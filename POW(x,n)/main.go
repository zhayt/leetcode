package main

import "math"

func helper(x float64, n int) float64 {
	// Base cases: If x is zero then 0^n (any number) would be zero
	// Calculate till the n is equal to one
	if x == 0 {
		return 0
	}
	if n == 0 {
		return 1
	}

	// Find the small fraction and multiply it with itself
	// ex: x = 2, n = 5; should be 2*2*2*2*2, but we need not to calculate all
	// if we could calculate 2*2=4 then multiply 4*4*2 = 32 which is equivalent to 2^5
	var res = helper(x, n/2)
	res = res * res

	// If the number is odd ie: 2^5, then only dividing by 2 will loose one 2
	// 2*2 * 2*2 = 16 where one 2 will be left
	// hence if we have a odd number, then we multiply res with one more x
	if n%2 != 0 {
		res = res * x
	}

	return res
}

func myPow(x float64, n int) float64 {
	var res = helper(x, int(math.Abs(float64(n))))

	// If n is negative then just do it 1/res, because (x^-n) is equivalent to (1/x^n)
	if n < 0 {
		return 1 / res
	}
	return res
}

func main() {
	myPow(2, 20)
}
