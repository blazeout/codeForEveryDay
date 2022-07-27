package _2_最小覆盖子串

import "math"

// mp1 和 mp2 我们只需要记录t中的字符的个数即可
func minWindow(s string, t string) (ans string) {
	if len(t) > len(s) {
		return ""
	}
	mp1, mp2 := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		mp2[t[i]]++
	}
	var check func() bool
	check = func() bool {
		for key, value := range mp2 {
			if mp1[key] < value {
				return false
			}
		}
		return true
	}
	left, right := 0, 0
	maxLength := math.MaxInt32
	for right < len(s) {
		// right字符如果碰到了 在t中的字符
		if mp2[s[right]] > 0 {
			mp1[s[right]]++
		}
		for left <= right && check() {
			if right-left+1 < maxLength {
				maxLength = right - left + 1
				ans = s[left : right+1]
			}
			if _, ok := mp2[s[left]]; ok {
				mp1[s[left]]--
			}
			left++
		}
		right++
	}
	return
}
