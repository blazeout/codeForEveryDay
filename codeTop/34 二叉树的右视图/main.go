package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) (res []int) {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	// 从右往左不就行了?
	for len(queue) > 0 {
		length := len(queue)
		length2 := length
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			if length == length2 {
				res = append(res, node.Val)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			length--
		}
	}
	return res
}

func main() {
	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4}
	node2 := &TreeNode{2, nil, node5}
	node3 := &TreeNode{Val: 3, Left: nil, Right: node4}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}
	res := rightSideView(node1)
	fmt.Println(res)
}
