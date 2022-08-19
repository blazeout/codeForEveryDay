package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{0, head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		node1 := cur.Next
		node2 := cur.Next.Next
		cur.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		cur = node1
	}
	return dummyHead.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = swapPairs(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
