package leetcode

func largestRectangleArea(heights []int) int {
	largest := 0
	for i := 0; i < len(heights); i++ {
		height := -1
		for j := i; j >= 0; j-- {
			if height == -1 || heights[j] < height {
				height = heights[j]
			}
			temp := (i - j + 1) * height
			if temp > largest {
				largest = temp
			}
			if (i+1)*height <= largest {
				break
			}
		}
	}
	return largest
}
