package _9_最长回文子串

func longestPalindrome(s string) string {
	// 中心扩散法
	// 以1个 或者2个位中心向外扩散
	result := ""
	for i := 0; i < len(s); i++ {
		left1, right1 := expand(s, i, i)
		left2, right2 := expand(s, i, i+1)
		if right1-left1+1 > len(result) {
			result = s[left1 : right1+1]
		}
		if right2-left2+1 > len(result) {
			result = s[left2 : right2+1]
		}
	}
	return result
}

func expand(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	return left + 1, right - 1
}
