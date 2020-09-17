package leetcode

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	r := [][]int{}
	if len(intervals) == 0 {
		return r
	}
	sort.SliceStable(intervals, func(p, q int) bool {
		return intervals[p][0] < intervals[q][0]
	})

	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right && intervals[i][1] >= right {
			right = intervals[i][1]
		} else if right < intervals[i][0] {
			r = append(r, []int{left, right})
			left = intervals[i][0]
			right = intervals[i][1]
		}
	}
	r = append(r, []int{left, right})
	return r
}
