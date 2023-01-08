package main

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param x int整型
 * @return bool布尔型
 */
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	var reverse int
	for x > reverse {
		reverse = reverse*10 + x%10
		x /= 10
	}
	if x == reverse || x == reverse/10 {
		return true
	}
	return false
}
