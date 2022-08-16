package main

import "fmt"

// Node is a node of a linked list.
type Node struct {
	Val        int
	Prev, Next *Node
}

// MyCircularDeque 使用链表来实现
type MyCircularDeque struct {
	// 虚拟头结点和尾节点
	head, rear  *Node
	length, cap int
}

func Constructor(k int) MyCircularDeque {
	head := &Node{0, nil, nil}
	tail := &Node{0, nil, nil}
	head.Next = tail
	tail.Prev = head
	return MyCircularDeque{
		head:   head,
		rear:   tail,
		length: 0,
		cap:    k,
	}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.length < d.cap {
		node := &Node{value, nil, nil}
		node.Next = d.head.Next
		node.Prev = d.head
		d.head.Next.Prev = node
		d.head.Next = node
		d.length++
		return true
	}
	return false
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.length < d.cap {
		node := &Node{value, nil, nil}
		node.Next = d.rear
		node.Prev = d.rear.Prev
		d.rear.Prev.Next = node
		d.rear.Prev = node
		return true
	}
	return false
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.head.Next != d.rear {
		d.head.Next = d.head.Next.Next
		d.head.Next.Prev = d.head
		return true
	}
	return false
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.rear.Prev != d.head {
		d.rear.Prev.Prev.Next = d.rear
		d.rear.Prev = d.rear.Prev.Prev
		return true
	}
	return false
}

func (d *MyCircularDeque) GetFront() int {
	if d.head.Next != d.rear {
		return d.head.Next.Val
	}
	return -1
}

func (d *MyCircularDeque) GetRear() int {
	if d.rear.Prev != d.head {
		return d.rear.Prev.Val
	}
	return -1
}

func (d *MyCircularDeque) IsEmpty() bool {
	if d.length == 0 {
		return true
	}
	return false
}

func (d *MyCircularDeque) IsFull() bool {
	return d.length == d.cap
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
func main() {
	obj := Constructor(4)
	front := obj.InsertFront(9)
	last := obj.DeleteLast()
	fmt.Println(front, last)
}
