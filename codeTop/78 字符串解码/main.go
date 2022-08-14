package main

import (
	"fmt"
	"strings"
)

func decodeString(s string) string {
	numStack := make([]int, 0)
	strStack := make([]string, 0) // 存储上一个str
	result := ""
	num := 0
	for i := range s {
		if s[i] <= '9' && s[i] >= '0' {
			num = num*10 + int(s[i]-'0')
		} else if s[i] == '[' {
			numStack = append(numStack, num)
			num = 0
			strStack = append(strStack, result)
			result = ""
		} else if s[i] == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			result = str + strings.Repeat(result, count)
		} else {
			result += string(s[i])
		}
	}
	return result
}

func main() {
	fmt.Println(decodeString("2[abc]3[cd]ef"))
}
