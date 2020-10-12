package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for k >= 0 {
		if j < 0 {
			break
		} else if i < 0 {
			nums1[k] = nums2[j]
			j--
		} else if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	// fmt.Printf("%v\n", nums1)
}
