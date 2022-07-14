package main

func search(nums []int, target int) int {
	// 二分查找, mid 分为三个区域
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[left] {
			if nums[mid] >= target && target >= nums[left] {
				// 位于A区
				right = mid
			} else {
				// 位于B区 或者C区
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 0, 1, 2}
	target := 3
	i := search(nums, target)
	println(i)
}
