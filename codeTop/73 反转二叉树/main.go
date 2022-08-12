package _3_反转二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		root.Left = right
		root.Right = left
		return root
	}
	return dfs(root)
}
