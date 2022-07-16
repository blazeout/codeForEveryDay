package _1_相交链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	listA, listB := headA, headB
	for listA != nil || listB != nil {
		if listA == nil {
			listA = headB
		}
		if listB == nil {
			listB = headA
		}
		if listA == listB {
			return listA
		}
		listA = listA.Next
		listB = listB.Next
	}
	return nil
}
