package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ans := nums[0] + nums[1] + nums[len(nums)-1]
	diff := int(math.Abs(float64(ans - target)))
	for i := 1; i < len(nums)-2; i++ {
		for low, high := i+1, len(nums)-1; low < high; {

			if nums[i]+nums[low]+nums[high] == target {
				return target
			} else if int(math.Abs(float64(nums[i]+nums[low]+nums[high]-target))) < diff {
				diff = int(math.Abs(math.Abs(float64(nums[i]+nums[low]+nums[high]) - math.Abs(float64(target)))))
				ans = nums[i] + nums[low] + nums[high]
			}

			if nums[i]+nums[low]+nums[high] > target {
				high--
			} else {
				low++
			}
		}
	}
	return ans
}

//func threeSumClosest(nums []int, target int) int {
//	sort.Ints(nums)
//	var ans int
//	diff := math.MaxInt
//
//	for i := 1; i < len(nums)-2; i++ {
//		low := i + 1
//		high := len(nums) - 1
//
//		for low < high {
//			if nums[i]+nums[low]+nums[high] == target {
//				ans = target
//				return ans
//			} else if int(math.Abs(float64(nums[i]+nums[low]+nums[high]-target))) <= diff {
//				diff = int(math.Abs(float64(nums[i] + nums[low] + nums[high] - target)))
//				ans = nums[i] + nums[low] + nums[high]
//			}
//
//			if nums[i]+nums[low]+nums[high] > target {
//				high--
//			} else {
//				low++
//			}
//		}
//	}
//	return ans
//}

// 1 2 3 4 5 6 7 8 9
func main() {
	fmt.Println(threeSumClosest([]int{1, 1, 1, 0}, 100))
}
