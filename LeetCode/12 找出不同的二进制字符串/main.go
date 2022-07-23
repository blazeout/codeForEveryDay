package _2_找出不同的二进制字符串

import (
	"strconv"
)

func findDifferentBinaryString(nums []string) string {
	mp := map[string]bool{}
	for _, str := range nums {
		mp[str] = true
	}
	flag := false
	ans := ""
	n := len(nums)
	var dfs func(index int, str string)
	dfs = func(index int, str string) {
		if index == n {
			if mp[str] {
				return
			}
			ans = str
			flag = true
			return
		}
		// 每一位可以选择
		for i := 0; i <= 1; i++ {
			dfs(index+1, str+strconv.Itoa(i))
			if flag {
				return
			}
		}
	}
	dfs(0, "")
	return ans
}
