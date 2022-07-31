package _6_最大层内元素和

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	curLevel := math.MaxInt32
	maxSum := math.MinInt32
	level := 1
	for len(queue) > 0 {
		length := len(queue)
		sum := 0
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			length--
		}
		if sum > maxSum {
			maxSum = sum
			curLevel = level
		}
		level++
	}
	return curLevel
}
