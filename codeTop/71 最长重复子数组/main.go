package _1_最长重复子数组

func findLength(nums1 []int, nums2 []int) int {
	// 动态规划, dp[1][5] 代表 A[:1] 和 B[:5] 中的最大的公共子数组长度
	// 边界条件 dp[0][1] 代表空数组 与 B[:1]的最大长度就是0
	dp := make([][]int, len(nums1)+1)
	for i := range dp {
		dp[i] = make([]int, len(nums2)+1)
	}
	maxRes := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				maxRes = max(maxRes, dp[i][j])
			}
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
