package leetcode

func minPathSum(grid [][]int) int {
	r := make([][]int, len(grid))
	for i := 0; i < len(r); i++ {
		r[i] = make([]int, len(grid[0]))
	}
	r[0][0] = grid[0][0]
	for i := 1; i < len(grid[0]); i++ {
		r[0][i] = r[0][i-1] + grid[0][i]
	}
	for i := 1; i < len(grid); i++ {
		r[i][0] = r[i-1][0] + grid[i][0]
	}

	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			if r[i-1][j] < r[i][j-1] {
				r[i][j] = r[i-1][j] + grid[i][j]
			} else {
				r[i][j] = r[i][j-1] + grid[i][j]
			}
		}
	}
	return r[len(grid)-1][len(grid[0])-1]
}
