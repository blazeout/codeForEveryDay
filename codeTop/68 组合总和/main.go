package _8_组合总和

func combinationSum(candidates []int, target int) [][]int {
	// 数组里面的每一个数 我们都可以是选择或者不选择
	// 每一个数我们都可以使用多次, 那么就要在dfs内部进行数组循环
	res := [][]int{}
	var dfs func(index, sum int, path []int)
	dfs = func(index, sum int, path []int) {
		if sum >= target {
			if sum == target {
				res = append(res, append([]int{}, path...))
			}
			return
		}
		for i := index; i < len(candidates); i++ {
			path = append(path, candidates[i])
			dfs(i, sum+candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0, []int{})
	return res
}
