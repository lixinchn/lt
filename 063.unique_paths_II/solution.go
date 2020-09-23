package leetcode

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	r := make([][]int, len(obstacleGrid))
	for i := 0; i < len(r); i++ {
		r[i] = make([]int, len(obstacleGrid[0]))
	}
	if obstacleGrid[0][0] == 0 {
		r[0][0] = 1
	}
	for i := 1; i < len(obstacleGrid[0]); i++ {
		if obstacleGrid[0][i] == 0 && obstacleGrid[0][i-1] == 0 {
			r[0][i] = r[0][i-1]
		}
	}
	for i := 1; i < len(obstacleGrid); i++ {
		if obstacleGrid[i][0] == 0 && obstacleGrid[i-1][0] == 0 {
			r[i][0] = r[i-1][0]
		}
	}

	for i := 1; i < len(obstacleGrid); i++ {
		for j := 1; j < len(obstacleGrid[0]); j++ {
			if obstacleGrid[i][j] == 1 {
				r[i][j] = 0
			} else {
				r[i][j] = r[i-1][j] + r[i][j-1]
			}
		}
	}
	return r[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}
