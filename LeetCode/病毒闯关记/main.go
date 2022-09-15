package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")
	atoi, _ := strconv.Atoi(split[0])
	length := atoi
	absss, _ := strconv.Atoi(split[1])
	scanner.Scan()
	arr := make([]int, length)
	split = strings.Split(scanner.Text(), " ")
	maxV := math.MinInt32
	minV := math.MaxInt32
	for i := 0; i < length; i++ {
		atoi, _ = strconv.Atoi(split[i])
		arr[i] = atoi
		maxV = max(maxV, arr[i])
		minV = min(minV, arr[i])
	}
	// 定义一个初始值
	count := 0
	first := minV
	maxNow := 0
	for ; first <= maxV; first++ {
		for i := 0; i < length; i++ {
			if abs(arr[i]-first) <= absss {
				count++
			}
		}

		if count > maxNow {
			maxNow = count
		}
		count = 0
	}

	fmt.Println(length - maxNow)
}
func min(v int, i int) int {
	if v < i {
		return v
	}
	return i
}

func max(v int, i int) int {
	if v > i {
		return v
	}
	return i
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
