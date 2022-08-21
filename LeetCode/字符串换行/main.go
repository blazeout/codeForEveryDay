package main

import (
	"fmt"
	"strings"
)

const length = 16

func getLastLine(str string, width int) string {
	sum := 0
	str = strings.TrimSpace(str)
	arr := strings.Split(str, " ")
	arr2 := []string{}
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			arr2 = append(arr2, arr[i])
		}
		if arr[i] == "" {

			// 向后找到第一个不为空格的地方
			j := i
			for arr[j] == "" {
				j++
			}
			i = j - 1
		}
	}
	arr = arr2
	for i := 0; i < len(arr); i++ {
		sum = len(arr[i]) * length
		if sum > width {
			return "error"
		}
	}
	return arr[len(arr)-1]
}

func main() {
	//line := getLastLine("     a", 18)
	//fmt.Println(line)
	//lastLine := getLastLine("abcdefg", 1000)
	//fmt.Println(lastLine)
	s := getLastLine("abcd         efg", 68)
	fmt.Println(s)
}
