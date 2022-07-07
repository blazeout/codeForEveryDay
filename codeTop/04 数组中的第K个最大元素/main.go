package main

import (
	"container/heap"
	"fmt"
)

// 使用了官方包的堆 接下来要使用堆排序
func findKthLargest(nums []int, k int) int {
	// 堆排序 对部分排序 大顶堆是排升序 小顶堆排降序 所以我们需要大顶堆
	temp := IntHeap(nums)
	//copy(temp, nums)
	h := &temp
	heap.Init(h)
	ret := 0
	for i := 0; i < k; i++ {
		ret = heap.Pop(h).(int)
	}
	return ret
}

type IntHeap []int

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h IntHeap) Len() int {
	return len(h)
}

func (h *IntHeap) Push(v interface{}) {
	*h = append(*h, v.(int))
}

func (h *IntHeap) Pop() (v interface{}) {
	old := *h
	length := len(*h)
	x := old[length-1]
	old = old[:length-1]
	*h = old
	return x
}

func findKthLargestOfHeapSort(nums []int, k int) int {
	// 使用堆排序, 排他个K次即可
	size := len(nums)
	length := size
	createHeap(nums, size)
	for i := size - 1; i > size-k; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		length--
		shiftDown(nums, 0, length)
	}
	return nums[0]
}

// 建立堆的函数
func createHeap(nums []int, size int) {
	// 从root 开始shiftDown
	root := (size - 2) / 2
	for ; root >= 0; root-- {
		shiftDown(nums, root, size)
	}
}

// 向下调整基本函数
func shiftDown(nums []int, parent int, size int) {
	child := 2*parent + 1
	for child < size {
		// 大顶堆
		if child+1 < size && nums[child+1] > nums[child] {
			child += 1
		}
		if nums[parent] < nums[child] {
			nums[child], nums[parent] = nums[parent], nums[child]
		} else {
			break
		}
		// 向下调整
		parent = child
		child = 2*parent + 1
	}
}

func main() {
	ret := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	fmt.Println(ret)
}
