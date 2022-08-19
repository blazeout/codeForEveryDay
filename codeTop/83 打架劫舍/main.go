package _3_打架劫舍

func rob(nums []int) int {
	i := len(nums)
	if i == 0 {
		return 0
	} else if i == 1 {
		return nums[0]
	}
	// 对于长度为2的数组我们在第二个地方也可以选择投前面一家或者 不偷 投自己的
	nums[1] = max(nums[0], nums[1])
	// 偷着家 + 投前两家 和 不偷着家 偷前一家进行选择
	// dp[i] 的含义是以i为结尾的 最大收益
	// dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	for i := 2; i < len(nums); i++ {
		nums[i] = max(nums[i-1], nums[i-2]+nums[i])
	}
	return nums[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
