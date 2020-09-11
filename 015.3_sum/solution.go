package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	m := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		low, high, preLow, preHigh := i+1, len(nums)-1, -1, -1
		for low < high {
			if preLow > -1 && nums[preLow] == nums[low] {
				low++
				continue
			}
			if preHigh > -1 && nums[preHigh] == nums[high] {
				high--
				continue
			}

			temp := nums[i] + nums[low] + nums[high]
			if temp == 0 {
				preLow = low
				preHigh = high
				m = append(m, []int{nums[i], nums[low], nums[high]})
				low++
				high--
			} else if temp < 0 {
				low++
			} else {
				high--
			}
		}
	}
	return m
}
