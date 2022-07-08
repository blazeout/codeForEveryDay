package main

import (
	"fmt"
	"math/rand"
)

// QuickSort 快速排序
func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, left, right int) {
	if left < right {
		mid := patition(nums, left, right)
		quickSort(nums, left, mid-1)
		quickSort(nums, mid+1, right)
	}
}

func patition(nums []int, left, right int) int {
	random := rand.Intn(right-left+1) + left
	nums[right], nums[random] = nums[random], nums[right]
	pivot := nums[right]
	l, r := left-1, left
	for ; r < right; r++ {
		if nums[r] <= pivot {
			l++
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[right], nums[l+1] = nums[l+1], nums[right]
	return l + 1
}

func main() {
	array := []int{2, 3, 4, 5, 1, 2, 3, 4, 6, 7, 5, 4, 3}
	QuickSort(array)
	fmt.Println(array)
}
