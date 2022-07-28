package _6_求根节点到叶节点的数字之和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) (sum int) {
	var dfs func(root *TreeNode, num int)
	dfs = func(root *TreeNode, num int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			sum += num*10 + root.Val
			return
		}
		num = num*10 + root.Val
		dfs(root.Left, num)
		dfs(root.Right, num)
	}
	dfs(root, 0)
	return
}
