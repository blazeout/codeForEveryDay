package _6_回文链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 断开 再 reverse 然后比较即可
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	newHead := slow.Next
	slow.Next = nil
	newHead = reverse(newHead)
	for head != nil && newHead != nil {
		if head.Val != newHead.Val {
			return false
		}
		head = head.Next
		newHead = newHead.Next
	}
	return true
}

func reverse(cur *ListNode) *ListNode {
	var pre *ListNode
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}
