package _3_字符串转整数

import "math"

// 先去除多于的空格, 如果此时index == len(s), 直接返回0
// 然后再判断正负号 记得i++, 如果没有则为正数
// 然後直接判斷下一個字符是否是数字, 如果是则继续, 如果不是則返回0

func myAtoi(s string) (result int) {
	if len(s) <= 0 {
		return
	}
	// 截取掉多于空白
	i := 0
	// 用来判断最后是正数还是负数
	flag := 1
	for i < len(s) && s[i] == ' ' {
		i++
	}
	if i == len(s) {
		return
	}
	s = s[i:]
	i = 0
	if s[i] == '-' {
		flag = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	// 不能循环到数字 如果 符号处理完的话 下一个还不是直接return 0
	for ; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			break
		}
		result = result*10 + int(s[i]-'0')
		if flag*result < math.MinInt32 {
			return math.MinInt32
		}
		if flag*result > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return result * flag
}
