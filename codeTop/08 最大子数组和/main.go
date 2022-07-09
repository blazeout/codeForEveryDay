package _8_最大子数组和

// 最大子数组和
func maxSubArray(nums []int) int {
	// write code here
	maxRes := nums[0]
	curRes := nums[0]
	for i := 1; i < len(nums); i++ {
		curRes = max(nums[i], curRes+nums[i])
		maxRes = max(maxRes, curRes)
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
