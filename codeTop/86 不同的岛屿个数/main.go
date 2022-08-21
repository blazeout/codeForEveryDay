package main

import "fmt"

func NumberOfDistinctIslands(grid [][]int) int {
	// 使用一个哈希Set
	// 对于每一个岛屿 我们是否有一种序列化的方式代表 一个岛屿的特殊标志
	// 就可以使用遍历规则 序列化为字符串 存入map中
	var dfs func(i, j int, path *string, dir string)
	m, n := len(grid), len(grid[0])
	count := 0
	mp := map[string]bool{}
	// 需要加上回去的路
	dfs = func(i, j int, path *string, dir string) {
		if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		*path += dir
		// 往下
		dfs(i+1, j, path, "D")
		// 往右
		dfs(i, j+1, path, "R")
		// 往上
		dfs(i-1, j, path, "U")
		// 往左
		dfs(i, j-1, path, "L")
		*path += "-" + dir
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				path := ""
				dfs(i, j, &path, "")
				if _, ok := mp[path]; !ok {
					count++
					mp[path] = true
				}
			}
		}
	}
	return count
}

func main() {
	island := [][]int{{1, 1, 0, 0, 1}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 1}, {0, 1, 0, 1, 1}}
	fmt.Println(NumberOfDistinctIslands(island))
}
