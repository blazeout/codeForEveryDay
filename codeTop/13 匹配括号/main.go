package _3_匹配括号

func isValid(s string) bool {
	// 使用栈, 碰到右括号就取出来
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != ')' && s[i] != ']' && s[i] != '}' {
			// 直接加入栈中
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}
			char := stack[len(stack)-1]
			if char == '(' && s[i] == ')' || char == '[' && s[i] == ']' || char == '{' && s[i] == '}' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
