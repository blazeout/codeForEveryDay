package _4_螺旋矩阵

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for {
		// 从left -> right
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}
		// 从top -> bottom
		for j := top; j <= bottom; j++ {
			ans = append(ans, matrix[j][right])
		}
		right--
		if right < left {
			break
		}
		// right -> left
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])

		}
		bottom--
		if bottom < top {
			break
		}
		// 从bottom -> top
		for j := bottom; j >= top; j-- {
			ans = append(ans, matrix[j][left])
		}
		left++
		if left > right {
			break
		}
	}
	return ans
}
