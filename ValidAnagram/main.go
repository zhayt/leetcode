package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	set := make(map[rune]int)
	for _, b := range s {
		set[b] += 1
	}
	for _, b := range t {
		if _, ok := set[b]; !ok {
			return false
			break
		}
		set[b] -= 1
	}
	for _, b := range t {
		if v := set[b]; v != 0 {
			return false
			break
		}
	}
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	chars := [26]int{}
	chart := [26]int{}
	for _, c := range s {
		chart[c-'a'] += 1
	}
	for _, c := range t {
		chart[c-'a'] += 1
	}
	for i := 0; i < 26; i++ {
		if chars[i] != chart[i] {
			return false
		}
	}
	return true
}
