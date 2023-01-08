package _1_reverseList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseListIterative(t *testing.T) {
	reverseArray := []int{5, 4, 3, 2, 1}
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	head = reverseListIterative(head)
	i := 0
	for head != nil {
		assert.Equal(t, reverseArray[i], head.Val, "they should be equal")
		head = head.Next
		i++
	}
}
