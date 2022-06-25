package _1_粉刷房子

import "math"

/**
author: wangjiahao.233
data: 2022/06/25
*/

func minCostRight(costs [][]int) int {
	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += min(costs[i-1][0], costs[i-1][1])
	}
	return min(costs[len(costs)-1][0], min(costs[len(costs)-1][1], costs[len(costs)-1][2]))
}

func minCostError(costs [][]int) int {
	// 错误解法 每次都选择最小的去刷了, 但其实 比如这次选择红 那么上次就只能刷蓝色或者绿色. 应该是有两种方案的
	minCostValue := 0
	colorArray := make([]int, len(costs))
	for i, cost := range costs {
		currentMinCost := math.MaxInt32
		for j, costMoney := range cost {
			if i == 0 && costMoney < currentMinCost {
				colorArray[i] = j
				currentMinCost = costMoney
			}
			if costMoney < currentMinCost && i >= 1 && colorArray[i-1] != j {
				colorArray[i] = j
				currentMinCost = costMoney
			}
		}
		minCostValue += currentMinCost
	}
	return minCostValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


