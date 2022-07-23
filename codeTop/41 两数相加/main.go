package _1_两数相加

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummyHead := &ListNode{0, nil}
	cur := dummyHead
	for l1 != nil || l2 != nil || carry != 0 {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
		carry = sum / 10
	}
	return dummyHead.Next
}
