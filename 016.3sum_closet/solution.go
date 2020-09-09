package leetcode

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	result := 0
	flag := 10000000
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		low, high := i+1, len(nums)-1
		for low < high {
			temp := nums[i] + nums[low] + nums[high]
			if temp >= target && temp-target < flag {
				result = temp
				flag = temp - target
			} else if temp < target && target-temp < flag {
				result = temp
				flag = target - temp
			}
			if result == target {
				return result
			}
			if temp >= target {
				high--
			} else {
				low++
			}
		}
	}
	return result
}
