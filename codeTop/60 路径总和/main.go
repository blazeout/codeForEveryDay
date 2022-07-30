package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	var dfs func(root *TreeNode, sum int, path []int)
	dfs = func(root *TreeNode, sum int, path []int) {
		if root == nil {
			return
		}
		path = append(path, root.Val)
		if root.Val == sum && root.Left == nil && root.Right == nil {
			res = append(res, append([]int{}, path...))
		}
		dfs(root.Left, sum-root.Val, path)
		dfs(root.Right, sum-root.Val, path)
	}
	dfs(root, targetSum, []int{})
	return res
}

// 不能发生拷贝, 使用unsafe.Pointer

func Str2Byte(str string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&str))
	b := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func Byte2Str(byte []byte) string {
	b := (*reflect.SliceHeader)(unsafe.Pointer(&byte))
	sh := reflect.StringHeader{
		Data: b.Data,
		Len:  b.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}

func main() {
	a := "aaa"
	fmt.Println(&a)
	ssh := (*reflect.StringHeader)(unsafe.Pointer(&a))
	b := (*[]byte)(unsafe.Pointer(ssh))
	fmt.Printf("%p\n", b)
	fmt.Printf("%p", b)
}
