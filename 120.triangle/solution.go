package leetcode

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][0]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][len(triangle[i-1])-1]
			} else {
				if triangle[i-1][j-1] >= triangle[i-1][j] {
					triangle[i][j] += triangle[i-1][j]
				} else {
					triangle[i][j] += triangle[i-1][j-1]
				}
			}
		}
	}
	ret := triangle[len(triangle)-1][0]
	for i := 1; i < len(triangle[len(triangle)-1]); i++ {
		if triangle[len(triangle)-1][i] < ret {
			ret = triangle[len(triangle)-1][i]
		}
	}
	return ret
}
