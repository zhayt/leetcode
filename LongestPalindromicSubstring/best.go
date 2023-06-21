package main

func LongestPalindrome(s_ string) string {
	s := getMatchString(s_)

	pArr := make([]int, len(s))
	C, R := -1, -1
	XC := 0

	for i := 0; i < len(s); i++ {
		pArr[i] = 1
		if R > i {
			pArr[i] = min(pArr[2*C-i], R-i)
		}

		for i-pArr[i] >= 0 && i+pArr[i] < len(s) && s[i-pArr[i]] == s[i+pArr[i]] {
			pArr[i]++
		}

		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}

		if pArr[i] > pArr[XC] {
			XC = i
		}
	}

	I, J := XC-pArr[XC]+1, XC+pArr[XC]-1
	result := make([]byte, 0, (J-I)/2)
	for i := I + 1; i < J; i += 2 {
		result = append(result, s[i])
	}

	return string(result)
}

func getMatchString(s string) []byte {
	result := make([]byte, 2*len(s)+1)

	for i := 0; i < len(s); i++ {
		result[2*i+1] = s[i]
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
