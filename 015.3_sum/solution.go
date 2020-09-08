package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	m := make([][]int, 0)
	t := make(map[int]bool)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if _, ok := t[nums[i]]; !ok {
			if r := twoSum(nums, nums[i]*-1, i, t); len(r) != 0 {
				m = append(m, r...)
			}
		}
	}
	return m
}

func twoSum(nums []int, target int, index int, t map[int]bool) [][]int {
	m := make(map[int]int)
	r := make([][]int, 0)
	prevAnother := 10000000
	for i := index + 1; i < len(nums); i++ {
		another := target - nums[i]
		if prevAnother != another {
			if _, ok := m[another]; ok {
				prevAnother = another
				r = append(r, []int{nums[i], target * -1, another})
				t[target*-1] = true
			}
		}
		m[nums[i]] = i
	}
	return r
}

// -4, -1, -1, 0, 1, 2
