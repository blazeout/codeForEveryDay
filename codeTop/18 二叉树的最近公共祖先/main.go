package _8_二叉树的最近公共祖先

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 递归写
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right == nil {
		return left
	}
	if left == nil && right != nil {
		return right
	}
	// 都为空就要返回空
	if left == nil && right == nil {
		return nil
	}
	// 如果两边都不为空那么就是自己 返回本身
	return root
}
