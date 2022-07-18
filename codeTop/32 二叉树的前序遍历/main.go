package _2_二叉树的前序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return
}
