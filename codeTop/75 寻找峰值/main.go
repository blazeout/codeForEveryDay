package main

import (
	"fmt"
	"math"
)

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)
	var test func(i int) int
	test = func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt64
		}
		return nums[i]
	}
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] > test(mid-1) && nums[mid] > test(mid+1) {
			return mid
		}
		// 如果是 > > > 那么 right = mid
		if test(mid) < test(mid+1) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return -1
}

func main() {
	nums := []int{3, 1, 2}
	peakElement := findPeakElement(nums)
	fmt.Println(peakElement)
}
