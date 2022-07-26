package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	strArr := []string{}
	i := 0
	for i < len(s) {
		j := i
		for j < len(s) && ((s[j] <= 'z' && s[j] >= 'a') || (s[j] <= 'Z' && s[j] >= 'A') || (s[j] <= '9' && s[j] >= '0')) {
			j++
		}
		if j != i {
			strArr = append(strArr, s[i:j])
		}
		// 把空格全部消除
		for j < len(s) && s[j] == ' ' {
			j++
		}
		i = j
	}
	for i, j := 0, len(strArr)-1; i < j; i, j = i+1, j-1 {
		strArr[i], strArr[j] = strArr[j], strArr[i]
	}
	return strings.Join(strArr, " ")
}

func main() {
	s := "EPY2giL"
	fmt.Println(reverseWords(s))
}
