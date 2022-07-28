package _5_二叉树的最大深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := dfs(root.Left), dfs(root.Right)
		return max(left, right) + 1
	}
	ret := dfs(root)
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
