package leetcode

import "fmt"

func rotate(matrix [][]int) {
	for i := 0; i <= len(matrix)/2; i++ {
		for j := i; j < len(matrix)-1-(i*2)+i; j++ {
			matrix[i][j], matrix[len(matrix)-j-1][i], matrix[len(matrix)-i-1][len(matrix)-j-1], matrix[j][len(matrix)-i-1] = matrix[len(matrix)-j-1][i], matrix[len(matrix)-i-1][len(matrix)-j-1], matrix[j][len(matrix)-i-1], matrix[i][j]
			// matrix[i][j] = matrix[len(matrix)-j-1][i]
			// matrix[len(matrix)-j-1][i] = matrix[len(matrix)-i-1][len(matrix)-j-1]
			// matrix[len(matrix)-i-1][len(matrix)-j-1] = matrix[j][len(matrix)-i-1]
			// matrix[j][len(matrix)-i-1] = matrix[i][j]
		}
	}
	fmt.Printf("%+v\n", matrix)
}
