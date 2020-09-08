package leetcode

func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		temp := min(height[i], height[j]) * (j - i)
		if max < temp {
			max = temp
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
