package _8_复原IP地址

import (
	"strconv"
	"strings"
)

/*
	important: !!!
*/

func restoreIpAddresses(s string) (res []string) {
	// dfs
	// 肯定不能以0开头
	var dfs func(SubStr []string, start int)
	dfs = func(SubStr []string, start int) {
		// 如果SubStr长度等于4 并且开始index刚好等于len(s) 那么就说明全部用完了
		if len(SubStr) == 4 && start == len(s) {
			str := strings.Join(SubStr, ".")
			res = append(res, str)
			return
		}
		if len(SubStr) == 4 && start != len(s) {
			return
		}
		// 接下来开始拼接字符串
		for length := 1; length <= 3; length++ {
			if start+length-1 >= len(s) {
				return
			}
			if length != 1 && s[start] == '0' {
				return
			}
			str := s[start : start+length]
			num, _ := strconv.Atoi(str)
			if num > 255 {
				return
			}
			SubStr = append(SubStr, str)
			dfs(SubStr, start+length)
			SubStr = SubStr[:len(SubStr)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}
