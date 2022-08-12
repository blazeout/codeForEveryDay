package _0_重新格式化字符串

func reformat(s string) string {
	arrOfDigit := make([]byte, 0)
	arrOfStr := make([]byte, 0)
	sumOfdigit, sumOfstr := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] <= 'z' && s[i] >= 'a' {
			sumOfstr++
			arrOfStr = append(arrOfStr, s[i])
		}
		if s[i] <= '9' && s[i] >= '0' {
			sumOfdigit++
			arrOfDigit = append(arrOfDigit, s[i])
		}
	}
	if abs(sumOfdigit-sumOfstr) > 1 {
		return ""
	}
	res := ""
	if sumOfdigit > sumOfstr {
		// 从数字先开始
		for i, j := 0, 0; i < len(arrOfDigit) || j < len(arrOfStr); i, j = i+1, j+1 {
			if i < len(arrOfDigit) {
				res += string(arrOfDigit[i])
			}
			if j < len(arrOfStr) {
				res += string(arrOfStr[i])
			}
		}
	} else {
		for i, j := 0, 0; i < len(arrOfDigit) || j < len(arrOfStr); i, j = i+1, j+1 {
			if j < len(arrOfStr) {
				res += string(arrOfStr[i])
			}
			if i < len(arrOfDigit) {
				res += string(arrOfDigit[i])
			}
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
