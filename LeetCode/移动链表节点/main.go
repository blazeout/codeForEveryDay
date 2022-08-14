package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := ""
	b := 0
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if len(a) == 0 {
		fmt.Println("")
	} else {
		arr := strings.Split(a, ",")
		// 建立链表
		atoi, _ := strconv.Atoi(arr[0])
		head := &ListNode{atoi, nil}
		cur := head
		for i := 1; i < len(arr); i++ {
			x, _ := strconv.Atoi(arr[i])
			cur.Next = &ListNode{x, nil}
			cur = cur.Next
		}
		node := rotateRight(head, b)
		for node != nil {
			if node.Next != nil {
				fmt.Print(node.Val)
				fmt.Print(",")
			} else {
				fmt.Print(node.Val)
			}
			node = node.Next
		}
	}

	//arr2 := make([]int, len(arr))
	//b = b % len(arr)
	//tempArr := make([]int, b)
	//for i := 0; i < len(tempArr); i++ {
	//	atoi, _ := strconv.Atoi(arr[len(arr)-b+i])
	//	tempArr[i] = atoi
	//}
	//for i := 0; i < len(arr); i++ {
	//	if i < b {
	//		arr2[i] = tempArr[i]
	//	} else {
	//		atoi, _ := strconv.Atoi(arr[i-b])
	//		arr2[i] = atoi
	//	}
	//}
	//for i := 0; i < len(arr2); i++ {
	//	if i == len(arr2)-1 {
	//		fmt.Print(arr2[i])
	//	} else {
	//		fmt.Print(arr2[i], ",")
	//	}
	//}
}

func rotateRight(head *ListNode, b int) *ListNode {
	if b == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
	}
	add := n - b%n
	if add == n {
		return head
	}
	cur.Next = head
	for i := 0; i < add; i++ {
		cur = cur.Next
	}
	ret := cur.Next
	cur.Next = nil
	return ret
}
