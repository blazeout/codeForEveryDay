package main

import (
	"fmt"
	"strconv"
)

func numDecoding(s string) int {
	// 0 只能与前面的字符结合, 如果0 找不到前面的数字与之结合 那么就无法解码 返回 0
	// 使用dfs 写吧
	charArr := map[string]bool{}
	for i := 1; i <= 26; i++ {
		charArr[strconv.Itoa(i)] = true
	}
	n := len(s)
	num := 0
	var dfs func(s string, index int)
	dfs = func(s string, index int) {
		if index == n {
			num++
			return
		}
		if _, ok := charArr[s[index:index+1]]; ok {
			dfs(s, index+1)
		}
		if index+1 < n {
			if _, ok := charArr[s[index:index+2]]; ok {
				dfs(s, index+2)
			}
		}
	}
	dfs(s, 0)
	return num
}

func main() {
	fmt.Println(numDecoding("12"))
}
