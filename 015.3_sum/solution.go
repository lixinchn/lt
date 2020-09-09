package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	m := [][]int{}
	t := map[int]bool{}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if _, ok := t[nums[i]]; !ok {
			mm := map[int]int{}
			prevAnother := 10000000
			for j := i + 1; j < len(nums); j++ {
				another := (nums[i] + nums[j]) * -1
				if prevAnother != another {
					if _, ok := mm[another]; ok {
						prevAnother = another
						m = append(m, []int{nums[j], nums[i], another})
						t[nums[i]] = true
					}
				}
				mm[nums[j]] = i
			}
		}
	}
	return m
}

// -4, -1, -1, 0, 1, 2
