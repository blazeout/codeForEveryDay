package _6_缺失的第一个正数

func firstMissingPositive(nums []int) int {
	// 如果等于1 就0号位置置为负数
	length := len(nums)
	for i := range nums {
		// 先把所有的负数变为 length+1
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}
	for i := range nums {
		if abs(nums[i]) <= length && nums[abs(nums[i])-1] > 0 {
			nums[abs(nums[i])-1] *= -1
		}
	}
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
