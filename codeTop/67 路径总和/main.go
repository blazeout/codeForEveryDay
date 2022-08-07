package _7_路径总和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	flag := false
	var dfs func(root *TreeNode, target int)
	dfs = func(root *TreeNode, target int) {
		if root == nil {
			return
		}
		if root.Val == target && root.Left == nil && root.Right == nil {
			flag = true
		}
		dfs(root.Left, target-root.Val)
		dfs(root.Right, target-root.Val)
	}
	dfs(root, targetSum)
	if flag {
		return true
	}
	return false
}
