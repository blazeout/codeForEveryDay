package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calc(str, fullStr string) int {
	res := math.MaxInt32
	mp1 := make(map[byte]int)
	mp2 := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		mp1[str[i]]++
	}
	for i := 0; i < len(fullStr); i++ {
		mp2[fullStr[i]]++
	}
	for i, v := range mp1 {
		res = min(res, mp2[i]/v)
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	times := 0
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	times, _ = strconv.Atoi(input.Text())

	for i := 0; i < times; i++ {
		res := 0
		str := ""
		input.Scan()
		str = input.Text()
		length := 0
		input.Scan()
		length, _ = strconv.Atoi(input.Text())
		strArr := make([]string, length)
		input.Scan()
		str1 := input.Text()
		strArr = strings.Split(str1, " ")
		for i2 := 0; i < length; i++ {
			res = max(res, calc(strArr[i2], str))
		}
		fmt.Println(res)
	}
}
