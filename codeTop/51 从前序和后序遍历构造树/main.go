package main

import "fmt"

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

type Student struct {
	Name string
	Age  int
}

func main() {
	student := []Student{
		{"张三", 18},
		{"李四", 20},
		{"王五", 22},
	}
	mp := make(map[string]*Student)
	// range 生成的v 是从student中取出的值, &v的话就是一直是v的地址, 后面改变了v的值
	// 那么map中的value也会改变
	for _, v := range student {
		mp[v.Name] = &v
	}
	fmt.Println(mp)
}
