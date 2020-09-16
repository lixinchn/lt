package leetcode

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return result
	}
	up, down, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	for {
		for i := left; i <= right; i++ {
			result = append(result, matrix[up][i])
		}
		up++
		if up > down {
			break
		}

		for i := up; i <= down; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		for i := right; i >= left; i-- {
			result = append(result, matrix[down][i])
		}
		down--
		if down < up {
			break
		}

		for i := down; i >= up; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return result
}
