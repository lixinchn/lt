package leetcode

func uniquePaths(m int, n int) int {
	r := make([][]int, m)
	for i := 0; i < m; i++ {
		r[i] = make([]int, n)
		r[i][0] = 1
	}
	for i := 0; i < n; i++ {
		r[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			r[i][j] = r[i-1][j] + r[i][j-1]
		}
	}
	return r[m-1][n-1]
}
