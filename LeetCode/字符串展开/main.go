package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := ""
	newStr := ""
	fmt.Scan(&str)
	strArr := []string{}
	temp := string(str[0])
	for i := 1; i < len(str); i++ {
		if str[i] != '(' {
			temp += string(str[i])
		} else {
			// 从后面开始找到第一个')'
			num := ""
			for j := i + 1; j < len(str); j++ {
				if str[j] != ')' {
					num += string(str[j])
				} else {
					strArr = append(strArr, temp, num)
					i = j
					temp = ""
					break
				}
			}
		}
	}
	for i := 0; i < len(strArr); i += 2 {
		atoi, _ := strconv.Atoi(strArr[i+1])
		newStr += strings.Repeat(strArr[i], atoi)
	}
	fmt.Print(newStr)
}
