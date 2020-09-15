package leetcode

func trap(height []int) int {
	valume := 0
	xPool, yPool := []int{}, []int{}
	for i := 0; i < len(height); i++ {
		if height[i] == 0 {
			continue
		}
		if len(xPool) == 0 {
			xPool = append(xPool, i)
			yPool = append(yPool, height[i])
			continue
		}

		baseHeight := 0
		for len(yPool) > 0 && height[i] >= yPool[len(yPool)-1] {
			x, y := xPool[len(xPool)-1], yPool[len(yPool)-1]
			xPool, yPool = xPool[:len(xPool)-1], yPool[:len(yPool)-1]
			valume += (y - baseHeight) * (i - x - 1)
			if y > baseHeight {
				baseHeight = y
			}
		}
		if len(yPool) > 0 && height[i] < yPool[len(yPool)-1] && height[i] > baseHeight {
			valume += (height[i] - baseHeight) * (i - xPool[len(xPool)-1] - 1)
		}
		xPool = append(xPool, i)
		yPool = append(yPool, height[i])
	}
	return valume
}
