package _3_二叉树的中序遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历是把自己压入栈中, 而前序遍历是把root.Right压入栈中

func inorderTraversal(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}
	return
}
