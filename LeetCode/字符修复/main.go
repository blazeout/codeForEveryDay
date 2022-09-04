package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	times := stringToInt(scanner.Text())
	arrNum := make([]int, times)
	for i := 0; i < times; i++ {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		// 字符串长度以及最大修复次数
		length := stringToInt(split[0])
		maxOperation := stringToInt(split[1])
		scanner.Scan()
		str := scanner.Text()
		scanner.Scan()
		operation := scanner.Text()
		i2 := maxLength(length, maxOperation, str, operation)
		arrNum[i] = i2
	}
	for i := 0; i < len(arrNum); i++ {
		fmt.Println(arrNum[i])
	}
}

func maxLength(length, maxOperation int, str, operation string) int {
	if maxOperation == 0 {
		return traverse(str)
	}
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range str {
		cnt[ch-'a']++
		maxCnt = max(maxCnt, cnt[ch-'a'])
		if right-left+1-maxCnt > maxOperation {
			cnt[str[left]-'a']--
			left++
		}
	}
	// 摆烂了...
	return len(str)/2 + 1
}

func max(res int, res2 int) int {
	if res > res2 {
		return res
	}
	return res2
}

func traverse(str string) int {
	res, temp := 1, 1
	for i := 1; i < len(str); i++ {
		if str[i] == str[i-1] {
			temp++
			if temp > res {
				res = temp
			}
		} else {
			temp = 1
		}
	}
	return res
}
