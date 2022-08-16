package main

import "fmt"

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mp := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		mp[nums[i]] = true
	}
	maxRes := 1
	for i := 0; i < len(nums); i++ {
		// 确保我们从序列的第一个开始遍历 ！！！
		if !mp[nums[i]-1] {
			count := 1
			temp := nums[i]
			for mp[temp+1] {
				count++
				temp++
			}
			maxRes = max(maxRes, count)
		}
	}
	return maxRes
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
