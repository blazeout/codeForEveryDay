package main

import "fmt"

func calcMinValue() {
	n, m := 6, 2

	arr := []int{7, 8, 1, 4, 3, 2}

	slidingWindow := []int{}

	for i, j := 1-m, 0; j < n; j, i = j+1, i+1 {
		if j == 0 {
			fmt.Println(0)
		} else {
			// 去除超出界限的值
			if i >= 2 && slidingWindow[0] == arr[i-2] {
				slidingWindow = slidingWindow[1:]
			}
			// 保持单调递增栈
			for len(slidingWindow) > 0 && slidingWindow[len(slidingWindow)-1] > arr[j-1] {
				slidingWindow = slidingWindow[:len(slidingWindow)-1]
			}
			slidingWindow = append(slidingWindow, arr[j-1])
			if j > 0 {
				fmt.Println(slidingWindow[0])
			}
		}
	}
}

// 6 2
// 7 8 1 4 3 2

func main() {
	calcMinValue()
}
