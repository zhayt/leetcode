package main

import "fmt"

func reverse(x int) int {
	result := 0
	for x >= 10 || x <= -10 {
		result *= 10
		d := x % 10
		if d < 0 {
			d *= -1
		}
		result += d
		x /= 10
	}
	if result*10 < -2147483648 || result*10 > 2147483647 {
		return 0
	}
	if x < 0 {
		return (result*10 + x*-1) * -1
	} else {
		return result*10 + x
	}
}

func main() {
	fmt.Println(reverse(100))
}
