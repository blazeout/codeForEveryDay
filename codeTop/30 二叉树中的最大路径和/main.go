package _0_二叉树中的最大路径和

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxRes := math.MinInt32
	// 当一条路的总贡献比0还小就返回0, 不要返回负值给我了
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		sum := left + right + root.Val
		maxRes = max(maxRes, sum)
		outPut := root.Val + max(left, right)
		return max(0, outPut)
	}
	dfs(root)
	return maxRes
}


func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}
