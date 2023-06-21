package main

func LengthOfLongestSubstring(s string) (max int) {
	for i, _ := range s {
		var letter [95]bool
		count := 0
		for j := i; j < len(s); j++ {
			if s[j] >= ' ' && s[j] <= '~' {
				if !letter[s[j]-' '] {
					count++
					letter[s[j]-' '] = true
				} else {
					break
				}
			}
		}
		if max < count {
			max = count
		}
	}
	return
}
