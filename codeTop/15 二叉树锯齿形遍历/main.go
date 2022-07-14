package _5_二叉树锯齿形遍历

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	ans := [][]int{}
	for len(queue) > 0 {
		length := len(queue)
		temp := []int{}
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			length--
		}
		ans = append(ans, temp)
	}
	for i := 0; i < len(ans); i++ {
		if i%2 == 1 {
			for j := 0; j < len(ans[i])/2; j++ {
				ans[i][j], ans[i][len(ans[i])-1-j] = ans[i][len(ans[i])-1-j], ans[i][j]
			}
		}
	}
	return ans
}
