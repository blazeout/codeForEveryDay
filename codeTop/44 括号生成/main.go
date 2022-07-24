package _4_括号生成


func generateParenthesis(n int) (ans []string) {
	// 只有当右括号的数量大于左括号的数量时 才可以选择右括号
	var dfs func(str string, left, right int)
	dfs = func(str string, left, right int) {
		if left == 0 && right == 0 {
			ans = append(ans, str)
			return
		}
		if left > 0 {
			dfs(str + "(", left-1, right)
		}
		if left < right {
			dfs(str + ")", left, right-1)
		}
	}
	dfs("", n, n)
	return ans
}
