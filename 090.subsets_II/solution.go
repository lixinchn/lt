package leetcode

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	flag := map[string]bool{}
	results := [][]int{[]int{}}
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		do(nums, i, []int{}, flag, "", &results)
	}
	return results
}

func do(nums []int, i int, temp []int, flag map[string]bool, flagPoint string, results *[][]int) {
	if _, ok := flag[fmt.Sprintf("%s%d", flagPoint, nums[i])]; !ok {
		flag[fmt.Sprintf("%s%d", flagPoint, nums[i])] = true
		temp = append(temp, nums[i])
		dest := make([]int, len(temp))
		copy(dest, temp)
		*results = append(*results, dest)
	}
	for j := i + 1; j < len(nums); j++ {
		do(nums, j, temp, flag, fmt.Sprintf("%s%d", flagPoint, nums[i]), results)
	}
}
