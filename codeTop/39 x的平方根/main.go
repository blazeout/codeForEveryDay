package _9_x的平方根

// 二分查找 缩小区间

func mySqrt(x int) int {
	ans := 0
	l, r := 0, x
	for l <= r {
		mid := (r-l)/2 + l
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
