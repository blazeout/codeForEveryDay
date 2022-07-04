package _6_最小绝对差

import (
	"math"
	"sort"
)

// 先排序再搜索最小绝对差

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	minValue := math.MaxInt32
	ans := [][]int{}
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < minValue {
			ans = [][]int{}
			ans = append(ans, []int{arr[i-1], arr[i]})
			minValue = arr[i] - arr[i-1]
		} else if arr[i]-arr[i-1] == minValue {
			ans = append(ans, []int{arr[i-1], arr[i]})
		}
	}
	return ans
}
