package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{0, head}
	curHead := dummyHead
	tail := dummyHead
	for curHead.Next != nil {
		times := k
		for times > 0 {
			tail = tail.Next
			times--
			if tail == nil {
				return head
			}
		}
		tempNext := tail.Next
		tempHead := curHead.Next
		curHead.Next = nil
		tail.Next = nil
		newHead, newTail := reverse(tempHead, tail)
		curHead.Next = newHead
		newTail.Next = tempNext
		curHead = newTail
		tail = newTail
	}
	return dummyHead.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return tail, head
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n5 := &ListNode{5, nil}
	n1.Next = n2
	n2.Next = nil
	n3.Next = n4
	n4.Next = n5
	reverseKGroup(n1, 2)
}
