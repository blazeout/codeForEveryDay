package main

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func DagPathNum(nodes [][]int) int {
	// dfs 深度遍历
	count := 0
	rooted := &Node{0, nil}
	mp := make(map[int]*Node)
	for i := 0; i < len(nodes); i++ {
		temp := getChildren(nodes[i], mp)
		var node *Node
		if mp[i] == nil {
			node = &Node{Val: i, Children: temp}
		} else {
			node = mp[i]
			node.Children = temp
		}
		if i == 0 {
			rooted = node
		}
		mp[i] = node
	}
	// 从0开始遍历
	isVisited := map[*Node]bool{}
	length := len(nodes)
	var dfs func(root *Node)
	dfs = func(root *Node) {
		if root == nil || isVisited[root] {
			return
		}
		if root.Val == length {
			count += 1
			return
		}
		for i := 0; i < len(root.Children); i++ {
			isVisited[root] = true
			dfs(root.Children[i])
			delete(isVisited, root)
		}
		return
	}
	dfs(rooted)
	return count
}

func getChildren(ints []int, mp map[int]*Node) (temp []*Node) {
	for i := 0; i < len(ints); i++ {
		if mp[ints[i]] != nil {
			temp = append(temp, mp[ints[i]])
		} else {
			mp[ints[i]] = &Node{Val: ints[i]}
			temp = append(temp, mp[ints[i]])
		}
	}
	return
}

func main() {
	arr := [][]int{{1, 2, 3}, {3}, {3}, {4}, {}}
	num := DagPathNum(arr)
	fmt.Println(num)
}
