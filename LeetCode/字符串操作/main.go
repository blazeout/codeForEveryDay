package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 返回满足题意的最小操作数
 * @param str string字符串 给定字符串
 * @return int整型
 */

func calc(n int) int {
	sum := 0
	for n >= 2 {
		sum += n / 2
		n = n/2 + n%2
	}
	return sum
}

func minOperations(str string) int {
	cnt := [26]int{}
	for _, ch := range str {
		cnt[ch-'a']++
	}
	sum := 0
	for i := 0; i < len(cnt); i++ {
		sum += calc(cnt[i])
	}
	return sum
}

func main() {
	operations := minOperations("abab")
	fmt.Println(operations)
}
