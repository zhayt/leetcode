package main

func maxConsecutiveAnswers(answerKey string, k int) int {
	if k >= len(answerKey) {
		return k
	}

	return max(slidingWindow(answerKey, k, 'T'), slidingWindow(answerKey, k, 'F'))
}

func slidingWindow(answerKey string, k int, target rune) int {
	left := 0
	resp := 0
	count := 0

	for right, item := range answerKey {
		if rune(item) != target {
			count++
		}
		for count > k {
			if rune(answerKey[left]) != target {
				count--
			}
			left++
		}
		resp = max(resp, right-left+1)
	}
	return resp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	maxConsecutiveAnswers("TTFTTFTT", 3)
}
