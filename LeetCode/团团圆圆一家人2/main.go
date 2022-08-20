package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getFamily(root *TreeNode) int {
	// 尽量从下往上遍历
	count := 0
	mp := make(map[*TreeNode]bool)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		if root.Left != nil && root.Right != nil && !mp[root.Left] && !mp[root.Right] {
			mp[root.Left] = true
			mp[root.Right] = true
			mp[root] = true
			count++
		}
	}
	dfs(root)
	return count
}

func main() {
	root1 := &TreeNode{1, nil, nil}
	root2 := &TreeNode{2, nil, nil}
	root3 := &TreeNode{3, nil, nil}
	root4 := &TreeNode{4, nil, nil}
	root5 := &TreeNode{5, nil, nil}
	root6 := &TreeNode{6, nil, nil}
	root7 := &TreeNode{7, nil, nil}
	root1.Left = root2
	root1.Right = root3
	root2.Left = root4
	root2.Right = root5
	root3.Left = root6
	root3.Right = root7
	family := getFamily(root1)
	fmt.Println(family)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")
	fmt.Println(split)
}
