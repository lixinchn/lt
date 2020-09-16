package leetcode

func firstMissingPositive(nums []int) int {
	for j := 0; j < len(nums); j++ {
		for nums[j] < len(nums) && nums[j] > 0 && nums[nums[j]-1] != nums[j] {
			temp := nums[nums[j]-1]
			nums[nums[j]-1] = nums[j]
			nums[j] = temp
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
