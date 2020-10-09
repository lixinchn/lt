package leetcode

func largestRectangleArea(heights []int) int {
	largest := 0
	heights = append(heights, 0)
	stack := []int{}
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := 0
			if len(stack) == 0 {
				left = -1
			} else {
				left = stack[len(stack)-1]
			}
			area := (i - left - 1) * h
			if area > largest {
				largest = area
			}
		}
		stack = append(stack, i)
	}
	return largest
}
