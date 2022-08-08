package _9_多数元素

func majorityElement(nums []int) int {
	// 摩尔投票法
	num := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			count++
		} else {
			count--
		}
		if count < 0 {
			count = 1
			num = nums[i]
		}
	}
	return num
}
