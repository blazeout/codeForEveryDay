package _2_最大正方形

func maximalSquare(matrix [][]byte) int {
	// 暴力法
	// 只需要记录maxSide即可
	maxSide := 0
	m, n := len(matrix), len(matrix[0])
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				maxSide = max(maxSide, 1)
				// 从1开始扩充
				curMaxSide := min(m-i, n-j)
				for k := 1; k < curMaxSide; k++ {
					flag := true
					if matrix[i+k][j+k] == '0' {
						break
					}
					for m := 0; m < k; m++ {
						if matrix[i+k][j+m] == '0' || matrix[i+m][j+k] == '0' {
							flag = false
							break
						}
					}
					if flag {
						maxSide = max(maxSide, k+1)
					} else {
						break
					}
				}
			}
		}
	}
	return maxSide * maxSide
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
