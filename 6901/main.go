package main

import "fmt"

func distanceTraveled(mainTank int, additionalTank int) int {
	res := 0
	for mainTank-5 >= 0 && additionalTank-1 >= 0 {
		res += 50
		mainTank -= 4
		additionalTank--
	}

	res += mainTank * 10

	return res
}

func main() {
	fmt.Println(distanceTraveled(100, 0))
}
