package _8_最长的有效括号

func longestValidParentheses(s string) int {
	// 贪心的认为 当right > left时 不满足匹配 将left和right 都置为0
	left, right, maxLength := 0, 0, 0
	// 但是会遇到这种情况我们永远无法找出有效括号 ((() 因为此时left > right
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
		if right > left {
			left, right = 0, 0
		} else if left == right {
			maxLength = max(maxLength, 2*right)
		}
	}
	// 解决这种办法就是 从右往左再遍历一次
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else if s[i] == ')' {
			right++
		}
		if left > right {
			left, right = 0, 0
		} else if left == right {
			maxLength = max(maxLength, 2*left)
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
