package main

import "fmt"

/*
	https://leetcode.cn/contest/weekly-contest-299/problems/count-number-of-ways-to-place-houses/
*/

func countHousePlacements(n int) int {
	// 上排列 + 下排列
	// 先 加 全0的一种方案 和 一个坑一个房子的方案
	sum := 0
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			sum += C(n-i+1, i) * C(n-j+1, j)
		}
	}
	return sum % (1e9 + 7)
}

func C(n, m int) int {
	if n < m {
		return 0
	}
	ans := 1
	for i := 1; i <= (n - m); i++ {
		if n > m {
			ans *= n
			n--
		}
		ans /= i
	}
	return ans
}

func main() {
	c := C(3, 2)
	fmt.Println(c)
}
