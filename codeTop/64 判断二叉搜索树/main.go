package _4_判断二叉搜索树

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	flag := true
	before := math.MinInt64
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode)  {
		if root == nil || !flag  {
			return
		}
		// 二叉搜索树 中序遍历有序
		dfs(root.Left)
		// 只需要比较这次的比上次的值是否大即可
		if root.Val <= before {
			flag = false
			return
		}
		before = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return flag
}
