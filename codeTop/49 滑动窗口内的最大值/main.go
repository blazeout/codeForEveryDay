package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	res := make([]int, len(nums)-k+1)
	queue := make([]int, 0)
	for j, i := 0, 1-k; j < len(nums); j, i = j+1, i+1 {
		// 去除 窗口外的数字
		if i >= 1 && queue[0] == nums[i-1] {
			queue = queue[1:]
		}
		// 要保证queue 单调递减
		for len(queue) > 0 && queue[len(queue)-1] < nums[j] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[j])
		if i >= 0 {
			res[i] = queue[0]
		}
	}
	return res
}

func main() {
	ret := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Print(ret)
}
