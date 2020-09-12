package leetcode

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	for i := 0; i < len(candidates); i++ {
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

	for j := index; j < len(candidates); j++ {
		flag = append(flag, candidates[j])
		do(candidates, target-candidates[j], j, flag, result)
		flag = flag[:len(flag)-1]
	}
}

// 2, 3, 6, 7      7
