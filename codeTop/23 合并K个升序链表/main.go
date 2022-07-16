package _3_合并K个升序链表

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	n := len(lists)
	left := mergeKLists(lists[:n/2])
	right := mergeKLists(lists[n/2:])
	mergedList := merge(left, right)
	return mergedList
}

func merge(l1, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	cur := dummyHead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 == nil {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return dummyHead.Next
}
