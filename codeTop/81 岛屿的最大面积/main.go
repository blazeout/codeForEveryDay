package _1_岛屿的最大面积

func maxAreaOfIsland(grid [][]int) int {
	maxRes := 0
	dix := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
			return 0
		}
		sum := 1
		grid[i][j] = 0
		for _, dir := range dix {
			sum += dfs(i+dir[0], j+dir[1])
		}
		return sum
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				count := dfs(i, j)
				maxRes = max(maxRes, count)
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
