package leetcode

func subsets(nums []int) [][]int {
	results := [][]int{}
	do(nums, 0, []int{}, &results)
	return results
}

func do(nums []int, index int, temp []int, results *[][]int) {
	newTemp := make([]int, len(temp))
	copy(newTemp, temp)
	*results = append(*results, newTemp)
	if index == len(nums) {
		return
	}
	for ; index < len(nums); index++ {
		temp = append(temp, nums[index])
		do(nums, index+1, temp, results)
		temp = temp[:len(temp)-1]
	}
}
