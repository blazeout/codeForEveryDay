package _2_股票买卖的最大利润

func maxProfit(prices []int) (maxRes int) {
	minCost := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minCost {
			minCost = prices[i]
		}
		maxRes = max(maxRes, prices[i]-minCost)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
