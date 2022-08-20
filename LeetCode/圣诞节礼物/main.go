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
	n, _ := strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		count := 0
		// 糖果的 length 长度
		length := 0
		scanner.Scan()
		length, _ = strconv.Atoi(scanner.Text())
		if length == 1 {
			scanner.Scan()
			fmt.Println(count)
		} else {
			maxApple := 0
			maxCandy := 0
			scanner.Scan()
			splitArr1 := strings.Split(scanner.Text(), " ")
			scanner.Scan()
			splitArr2 := strings.Split(scanner.Text(), " ")
			arr := make([][]int, length)
			for i := 0; i < length; i++ {
				arr[i] = make([]int, 2)
				atoi, _ := strconv.Atoi(splitArr1[i])
				arr[i][0] = atoi
				i2, _ := strconv.Atoi(splitArr2[i])
				arr[i][1] = i2
				maxApple = max(maxApple, atoi)
				maxCandy = max(maxCandy, i2)
			}
			for i := 0; i < length; i++ {
				v1 := maxApple - arr[i][0]
				v2 := maxCandy - arr[i][1]
				if v1 > v2 {
					count += v2
					v1 -= v2
					arr[i][1] += v2
					arr[i][0] += v2
				} else {
					count += v1
					v2 -= v1
					arr[i][0] += v1
					arr[i][1] += v1
				}
				if arr[i][0] < maxApple {
					count += v1
				}
				if arr[i][1] < maxCandy {
					count += v2
				}
			}
			fmt.Println(count)
		}
	}
}

func max(apple int, i int) int {
	if apple > i {
		return apple
	}
	return i
}
