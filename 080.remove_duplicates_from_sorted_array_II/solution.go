package leetcode

func removeDuplicates(nums []int) int {
	i := 2
	for j := 2; j < len(nums); j++ {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
