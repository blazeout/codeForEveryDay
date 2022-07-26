package _1_从前序和后序遍历构造树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}
	head := preorder[0]
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == head {
			break
		}
	}
	root := &TreeNode{head, nil, nil}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}
