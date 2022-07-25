package _4_完全二叉树插入器

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type CBTInserter struct {
	Root  *TreeNode
	Queue []*TreeNode
}

func Constructor(root *TreeNode) CBTInserter {
	queue := []*TreeNode{root}
	var res []*TreeNode
	for len(queue) > 0 {
		length := len(queue)
		for length > 0 {
			node := queue[0]
			queue = queue[1:]
			res = append(res, node)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			length--
		}
	}
	return CBTInserter{
		Root:  root,
		Queue: res,
	}
}

func (i *CBTInserter) Insert(val int) int {
	node := &TreeNode{val, nil, nil}
	i.Queue = append(i.Queue, node)
	parentIndex := (len(i.Queue) - 2) / 2
	if i.Queue[parentIndex].Left == nil {
		i.Queue[parentIndex].Left = node
	} else {
		i.Queue[parentIndex].Right = node
	}
	return i.Queue[parentIndex].Val
}

func (i *CBTInserter) GetRoot() *TreeNode {
	return i.Root
}
