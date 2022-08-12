package _4_在排序数组中查找元素的第一个和最后一个位置

func searchRange(nums []int, target int) []int {
	// binarySearch
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	index1 := binarySearch(nums, target)
	index2 := binarySearch(nums, target+1)
	// 因为我们的target + 1 是可以不在数组中的
	if index1 == len(nums) || nums[index1] != target {
		return []int{-1, -1}
	}
	return []int{index1, index2 - 1}
}

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right-left)/2 + left
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
