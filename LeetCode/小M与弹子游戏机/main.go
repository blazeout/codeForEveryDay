package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// 只能选择最上面那一排的位置
func main() {
	// 高为 n 宽为 m
	n, m := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	splitNum := strings.Split(scanner.Text(), " ")
	n = stringToInt(splitNum[0])
	m = stringToInt(splitNum[1])
	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		for j := 0; j < m; j++ {
			arr[i][j] = stringToInt(split[j])
		}
	}
	// dfs返回此次遍历能过获取的价值最大值
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i < 0 || i >= n || j < 0 || j >= m {
			return 0
		}
		value := 0
		if arr[i][j] == -1 {
			// 直接选择左下 或者 右下
			value += max(dfs(i+1, j-1), dfs(i+1, j+1))
		} else {
			value = arr[i][j]
			value += dfs(i+1, j)
		}
		return value
	}
	// 棋盘已经构建好了
	maxRes := 0
	//for j := 0; j < m; j++ {
	//	res := dfs(0, j)
	//	maxRes = max(maxRes, res)
	//}
	maxRes = dfs(0, 0)
	print(maxRes)
}

func max(res int, res2 int) int {
	if res > res2 {
		return res
	}
	return res2
}

func stringToInt(s string) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}
