package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	root := &TreeNode{1, nil, nil}
	queue := []*TreeNode{root}
	for i := 1; i < n+1; i++ {
		left, right := 0, 0
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		left, _ = strconv.Atoi(split[0])
		right, _ = strconv.Atoi(split[1])
		if left == 0 {
			queue = append(queue, nil)
		} else {
			queue = append(queue, &TreeNode{left, nil, nil})
		}
		if right == 0 {
			queue = append(queue, nil)
		} else {
			queue = append(queue, &TreeNode{right, nil, nil})
		}
	}
	root = queue[0]
	for i := 0; i < len(queue); i++ {
		left := 2*i + 1
		right := 2*i + 2
		node := queue[i]
		if left < len(queue) && node != nil {
			node.Left = queue[left]
		}
		if right < len(queue) && node != nil {
			node.Right = queue[right]
		}
	}
	minValue := math.MaxInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		i := dfs(root.Left)
		i2 := dfs(root.Right)
		if i == -1 || i2 == -1 || abs(i-i2) > 1 {
			return -1
		}
		minValue = min(minValue, root.Val)
		return max(i, i2) + 1
	}
	dfs(root)
	fmt.Println(minValue)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func min(value int, val int) int {
	if value > val {
		return val
	}
	return value
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
