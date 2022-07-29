package _9_子集

/**
author: wangjiahao
https://leetcode.cn/problems/subsets/
*/

func subsets(nums []int) [][]int {
	result := [][]int{}
	// 对于每个元素我们都可以选择或者不选择两种方式
	var dfs func(path []int, j int)
	dfs = func(path []int, j int) {
		if j == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}
		dfs(path, j+1)
		path = append(path, nums[j])
		dfs(path, j+1)
		path = path[:len(path)-1]
	}
	dfs([]int{}, 0)
	return result
}
