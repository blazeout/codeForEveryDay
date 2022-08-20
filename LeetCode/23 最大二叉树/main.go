package _3_最大二叉树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	temp := 0
	maxRes := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > maxRes {
			maxRes = nums[i]
			temp = i
		}
	}
	return &TreeNode{nums[temp],
		constructMaximumBinaryTree(nums[:temp]),
		constructMaximumBinaryTree(nums[temp+1:])}
}
