package _9_逐步求和得到正数的最小值

func minStartValue(nums []int) int {
	sum := 0
	minRes := 0
	for i := range nums {
		sum += nums[i]
		if sum < minRes {
			minRes = sum
		}
	}
	return abs(minRes) + 1
}
func abs(minRes int) int {
	if minRes < 0 {
		return -minRes
	}
	return minRes
}
