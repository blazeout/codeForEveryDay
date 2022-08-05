package _7_在二叉树中增加一行

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	queue := []*TreeNode{root}
	level := 1
	flag := false
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			if level == depth-1 {
				temp := node.Left
				node.Left = &TreeNode{val, temp, nil}
				temp = node.Right
				node.Right = &TreeNode{val, nil, temp}
				flag = true
			} else {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			}
			length--
		}
		level++
		if flag {
			return root
		}
	}
	return root
}
