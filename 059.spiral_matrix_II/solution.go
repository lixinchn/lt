package leetcode

func generateMatrix(n int) [][]int {
	result := [][]int{}
	for i := 0; i < n; i++ {
		result = append(result, make([]int, n))
	}
	up, down, left, right := 0, n-1, 0, n-1
	x, y := 0, 0
	direct := 1
	for i := 1; i <= n*n; i++ {
		result[x][y] = i
		if direct == 1 {
			y++
		} else if direct == 2 {
			x++
		} else if direct == 3 {
			y--
		} else {
			x--
		}

		if y > right {
			up++
			direct = 2
			y--
			x++
		} else if x > down {
			right--
			direct = 3
			x--
			y--
		} else if y < left {
			down--
			direct = 4
			y++
			x--
		} else if x < up {
			left++
			direct = 1
			x++
			y++
		}
	}
	return result
}
