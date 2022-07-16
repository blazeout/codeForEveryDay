package _6_在指定位置反转链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{0, head}
	first, second := dummyHead, dummyHead
	for i := 0; i < left-1; i++ {
		first = first.Next
	}
	for i := 0; i < right; i++ {
		second = second.Next
	}
	reverseHead := first.Next
	tail := second.Next
	first.Next = nil
	second.Next = nil
	reverseHead, second = reverse(reverseHead)
	first.Next = reverseHead
	second.Next = tail
	return dummyHead.Next
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre, head
}
