package _9_重排链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	// 快慢指针
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	newHead := slow.Next
	slow.Next = nil
	newHead = reverseList(newHead)
	cur := head
	for cur != nil && newHead != nil {
		temp1 := cur.Next
		temp2 := newHead.Next
		cur.Next = newHead
		newHead.Next = temp1
		cur = temp1
		newHead = temp2
	}
	return
}

func reverseList(cur *ListNode) *ListNode {
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
