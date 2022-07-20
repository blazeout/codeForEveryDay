package _1_二维网格迁移

func shiftGrid(grid [][]int, k int) [][]int {
	// 二维转一维, 移动一位, 然后再转回来
	m, n := len(grid), len(grid[0])
	oneDimensional := make([]int, m*n)
	index := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			oneDimensional[index] = grid[i][j]
			index++
		}
	}
	k = k % (m * n)
	for i := 0; i < k; i++ {
		temp := oneDimensional[m*n-1]
		for j := len(oneDimensional) - 1; j >= 1; j-- {
			oneDimensional[j] = oneDimensional[j-1]
		}
		oneDimensional[0] = temp
	}
	index = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			grid[i][j] = oneDimensional[index]
			index++
		}
	}
	return grid
}
