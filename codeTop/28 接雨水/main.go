package _8_接雨水

func trap(height []int) int {
	// 使用两个数组保存该点的左右两边柱子的最大高度, 那么接的雨水就是 min(left, right) - height[i]
	leftMaxArr, rightMaxArr := make([]int, len(height)), make([]int, len(height))
	leftMax, rightMax := height[0], height[len(height)-1]
	leftMaxArr[0] = leftMax
	rightMaxArr[len(height)-1] = rightMax
	for i := 1; i < len(height); i++ {
		leftMaxArr[i] = max(leftMax, height[i])
		leftMax = max(leftMax, height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMaxArr[i] = max(height[i], rightMax)
		rightMax = max(rightMax, height[i])
	}
	sum := 0
	for i := 0; i < len(height); i++ {
		sum += min(leftMaxArr[i], rightMaxArr[i]) - height[i]
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
