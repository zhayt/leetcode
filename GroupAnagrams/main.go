package main

import (
	"fmt"
)

//
//func groupAnagrams(strs []string) [][]string {
//	var res [][]string
//	set := make(map[string]int)
//	index := 0
//	for _, s := range strs {
//		sorted := string(mergeSort(s))
//		if val, ok := set[sorted]; ok {
//			res[val] = append(res[val], s)
//			continue
//		}
//		set[sorted] = index
//		index++
//		res = append(res, []string{s})
//	}
//	return res
//}
//
//func mergeSort(items string) []byte {
//	if len(items) == 1 {
//		return []byte(items)
//	}
//	first := mergeSort(items[:len(items)/2])
//	second := mergeSort(items[len(items)/2:])
//	return merge(first, second)
//}
//
//func merge(a []byte, b []byte) []byte {
//	var merged []byte
//	i := 0
//	j := 0
//	for i < len(a) && j < len(b) {
//		n := bytes.Compare([]byte{a[i]}, []byte{b[j]})
//		if n == -1 {
//			merged = append(merged, a[i])
//			i++
//		} else {
//			merged = append(merged, b[j])
//			j++
//		}
//	}
//	for ; i < len(a); i++ {
//		merged = append(merged, a[i])
//	}
//	for ; j < len(b); j++ {
//		merged = append(merged, b[j])
//	}
//	return merged
//}

func groupAnagrams(strs []string) [][]string {
	res := make(map[[26]int][]string)
	for _, str := range strs {
		count := [26]int{}
		for _, s := range str {
			count[s-'a']++
		}
		res[count] = append(res[count], str)
	}
	a := make([][]string, len(res))
	for _, val := range res {
		a = append(a, val)
	}
	return a
}

func main() {
	a := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(a)
}
