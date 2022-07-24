package _5_最长公共子序列

func longestCommonSubsequence(text1 string, text2 string) int {
	// dp[i][j] 代表 text1[0:i-1] 和 text2[0:j-1] 之间的最长子序列
	// dp[0][j] == dp[i][0] 因为 都是空字符串和其他对比 所以肯定是0
	// return dp[m][n] 即可
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
