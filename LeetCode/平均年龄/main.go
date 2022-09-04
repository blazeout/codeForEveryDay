package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// 平均年龄, 计算N年后的平均年龄
func main() {
	W, Y, X, N := 0, 0, 0.0, 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")
	W, _ = strconv.Atoi(split[0])
	Y, _ = strconv.Atoi(split[1])
	X, _ = strconv.ParseFloat(split[2], 64)
	N, _ = strconv.Atoi(split[3])
	// 每次新招的人数
	number := float64(W) * X
	sum := W * Y
	remain := W
	for i := 0; i < N; i++ {
		if remain >= int(number) {
			sum -= int(number) * Y
			remain -= int(number)
		} else if remain < int(number) && remain > 0 {
			sum -= remain * Y
			sum -= (int(number) - remain) * 21
			remain = 0
		} else {
			sum -= int(number) * 21
		}
		sum += int(number) * 21
	}
	ceil := math.Ceil(float64(sum) / float64(W))
	fmt.Println(int(ceil))
}
