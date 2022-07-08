package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	// 先sort一下 这样就方便双指针查找
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		firstValue := nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			sum := firstValue + nums[left] + nums[right]
			if sum == 0 {
				ans = append(ans, []int{firstValue, nums[left], nums[right]})
				left++
				right--
				for left < len(nums) && nums[left] == nums[left-1] {
					left++
				}
				for right > 0 && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				// 小于0 left往右移
				left++
			} else {
				// 大于0 right往左移
				right--
			}
		}
	}
	return ans
}

func main() {
	sum := threeSum([]int{-1, 0, 1, 2, -1, -4})
	fmt.Println(sum)
}
