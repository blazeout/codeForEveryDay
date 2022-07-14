package _4_环形链表

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
环形链表
author: wangjiahao.233
*/

func hasCycle(head *ListNode) bool {
	// 快慢指针
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
