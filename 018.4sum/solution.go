package leetcode

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(nums)

	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// -5, -4, -2, -2, -2, -1, 0, 0, 1
			low, high, preLow, preHigh := j+1, len(nums)-1, j+1, len(nums)-1
			for low < high {
				if preLow != low {
					preValue := nums[preLow]
					preLow = low
					if preValue == nums[low] {
						low++
						continue
					}
				}
				if preHigh != high {
					preValue := nums[preHigh]
					preHigh = high
					if preValue == nums[high] {
						high--
						continue
					}
				}
				temp := nums[i] + nums[j] + nums[low] + nums[high]
				if temp == target {
					result = append(result, []int{nums[i], nums[j], nums[low], nums[high]})
					low++
					high--
				} else if temp > target {
					high--
				} else {
					low++
				}
			}
		}
	}
	return result
}
