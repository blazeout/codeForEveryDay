package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 字符串清理
func removeDuplicates(s string, k int) string {
	for {
		temp := 1
		flag := false
		for i := 1; i < len(s); i++ {
			if s[i] != s[i-1] {
				temp = 1
			} else {
				temp++
			}
			if temp == k {
				s = s[:i-k+1] + s[i+1:]
				break
			}
			if i == len(s)-1 {
				flag = true
			}
		}
		if flag {
			break
		}
	}
	return s
}

// 所有的公共祖先2, 找到所有的公共祖先并返回

func findCommonAncestors2(root *TreeNode, p *TreeNode, q *TreeNode) []*TreeNode {

	ans := []*TreeNode{}
	flag := false
	var dfs func(root, p, q *TreeNode) *TreeNode
	dfs = func(root, p, q *TreeNode) *TreeNode {
		if root == nil || root == p || root == q {
			return root
		}
		left := dfs(root.Left, p, q)
		right := dfs(root.Right, p, q)
		if left == nil && right != nil {
			if flag {
				ans = append(ans, root)
			}
			return right
		}
		if left != nil && right == nil {
			if flag {
				ans = append(ans, root)
			}
			return left
		}
		if left == nil && right == nil {
			return nil
		}
		flag = true
		if flag {
			ans = append(ans, root)
		}
		return root
	}
	dfs(root, p, q)
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return ans
}

func areNumbersAscending(s string) bool {
	splitArr := strings.Split(s, " ")
	before := -1
	for i := 0; i < len(splitArr); i++ {
		if splitArr[i][0] >= 'a' && splitArr[i][0] <= 'z' {
			continue
		}
		atoi, _ := strconv.Atoi(splitArr[i])
		if atoi <= 100 && atoi >= 2 {
			if before == -1 {
				before = atoi
			} else {
				if atoi <= before {
					return false
				}
				before = atoi
			}
		}
	}
	return true
}

func main() {
	//node9 := &TreeNode{6, nil, nil}
	//node8 := &TreeNode{Val: 11}
	//node7 := &TreeNode{8, nil, nil}
	//node6 := &TreeNode{0, nil, nil}
	//node5 := &TreeNode{2, node8, node9}
	//node4 := &TreeNode{3, nil, nil}
	//node3 := &TreeNode{1, node6, node7}
	//node2 := &TreeNode{5, node4, node5}
	//node1 := &TreeNode{9, node2, node3}
	//nodes := findCommonAncestors2(node1, node7, node6)
	//for _, node := range nodes {
	//	fmt.Println(node.Val)
	//}
	ascending := areNumbersAscending("sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s")
	fmt.Println(ascending)
}
