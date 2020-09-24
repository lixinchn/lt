package leetcode

func sortColors(nums []int) {
	for i, j, k := 0, 0, len(nums)-1; j <= k; j++ {
		if nums[j] == 0 {
			if i == j {
				continue
			}
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else if nums[j] == 1 {
			continue
		} else {
			nums[j], nums[k] = nums[k], nums[j]
			j--
			k--
		}
	}
}
