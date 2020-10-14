package leetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	res := [][]int{triangle[0]}
	for i := 1; i < len(triangle); i++ {
		line := make([]int, len(triangle[i]))
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				line[j] = res[i-1][0] + triangle[i][j]
			} else if j == len(triangle[i])-1 {
				line[j] = res[i-1][len(res[i-1])-1] + triangle[i][j]
			} else {
				if res[i-1][j-1] >= res[i-1][j] {
					line[j] = res[i-1][j] + triangle[i][j]
				} else {
					line[j] = res[i-1][j-1] + triangle[i][j]
				}
			}
		}
		res = append(res, line)
	}
	ret := res[len(res)-1][0]
	for i := 1; i < len(res[len(res)-1]); i++ {
		if res[len(res)-1][i] < ret {
			ret = res[len(res)-1][i]
		}
	}
	return ret
}
