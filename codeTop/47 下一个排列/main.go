package _7_下一个排列

func nextPermutation(nums []int) {
	// 找到第一个 nums[i] < nums[j]
	if len(nums) <= 1 {
		return
	}
	i, j, k := len(nums)-2, len(nums)-1, len(nums)-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		// 找到第一个nums[k] > nums[i]
		for nums[k] <= nums[i] {
			k--
		}
		nums[k], nums[i] = nums[i], nums[k]
	}
	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}