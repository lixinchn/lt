package leetcode

func maxSubArray(nums []int) int {
	max, temp := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if temp >= 0 {
			temp += nums[i]
		} else {
			if nums[i] > temp {
				temp = nums[i]
			}
		}

		if max < temp {
			max = temp
		}
	}
	return max
}
