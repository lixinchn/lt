package leetcocde

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(candidates)
	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}
		do(candidates, target-candidates[i], i, []int{candidates[i]}, &result)
	}
	return result
}

func do(candidates []int, target int, index int, flag []int, result *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		newFlag := make([]int, len(flag))
		copy(newFlag, flag)
		*result = append(*result, newFlag)
		return
	}

	for j := index + 1; j < len(candidates); j++ {
		if j > index+1 && candidates[j] == candidates[j-1] {
			continue
		}
		flag = append(flag, candidates[j])
		do(candidates, target-candidates[j], j, flag, result)
		flag = flag[:len(flag)-1]
	}
}

// 1, 2, 2, 2, 5
