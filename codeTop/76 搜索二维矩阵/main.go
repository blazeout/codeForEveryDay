package _6_搜索二维矩阵

func searchMatrix(matrix [][]int, target int) bool {
	i, j := 0, len(matrix[0])-1
	m, n := len(matrix), len(matrix[0])
	for i >= 0 && i < m && j >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
