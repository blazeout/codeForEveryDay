package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	root.Left = buildTree(preorder[1:i+1], inorder[:i])
	root.Right = buildTree(preorder[i+1:], inorder[i+1:])
	return root
}

func getBinaryTrees(preOrder []int, inOrder []int) []*TreeNode {
	// write code here
	ans := []*TreeNode{}
	i := 0
	for ; i < len(inOrder); i++ {
		if inOrder[i] == preOrder[0] {
			root := &TreeNode{preOrder[0], nil, nil}
			root.Left = buildTree(preOrder[1:i+1], inOrder[:i])
			root.Right = buildTree(preOrder[i+1:], inOrder[i+1:])
			ans = append(ans, root)
		}
	}
	return ans
}

func main() {

}
