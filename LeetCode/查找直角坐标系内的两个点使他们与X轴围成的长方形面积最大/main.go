package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	text = text[1:]
	text = text[:len(text)-1]
	split := strings.Split(text, ",")
	arr := make([]int, len(split))
	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(split[i])
	}
	// 两个点 x轴方向作为长方形的长, y轴方向作为长方形的宽
	// 贪心 从index 差距最大的两点开始遍历
	left, right := 0, len(arr)-1
	maxArea := 0
	for left < right {
		maxArea = max(maxArea, (right-left)*min(arr[right], arr[left]))
		left++
	}
	fmt.Println(maxArea)
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
