package main

import "fmt"

func main() {
	length := 0
	fmt.Scan(&length)
	result := length / 3
	remain := length % 3
	if length < 6 {
		result = 0
	} else if remain == 0 {
		result = 1
	} else if remain == 1 {
		// 有多少个空为n, 如果空为3个 那么 A3,1 * C26,1
		result = A(result+1, 1) * C(26, 1)
	} else if remain == 2 {
		result = A(result+1, 2) * C(26, 2)
	}
	fmt.Println(result % 1000000007)
}

// A3,2 == 6
func A(n, m int) int {
	result := 1
	for i := m; i > 0; i-- {
		result *= n
		n--
	}
	return result
}

// C3,2 == 3 C = A / m!
func C(n, m int) int {
	if m > (n - m) {
		m = n - m
	}
	son := A(n, m)
	mother := A(m, m)
	return son / mother
}
