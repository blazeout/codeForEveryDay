package _2_层数最深的叶子节点的和

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxLevel := 1
	maxSum := 0
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level == maxLevel {
			maxSum += root.Val
		}
		if level > maxLevel {
			maxLevel = level
			maxSum = root.Val
		}

		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 1)
	return maxSum
}
