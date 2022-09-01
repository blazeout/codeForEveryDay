package main

import (
	"fmt"
	"math"
)

//func splitNumber(n int) (int, int) {
//	// 对于每一个自然数我们求得了 最小的result时, 还可以将他两个 组成的数据给他拆分为最小的
//}

func main() {
	mp := map[int]int{
		6:  5,
		12: 7,
	}
	number := 0
	fmt.Scan(&number)
	result := math.MaxInt64
	if number == 2 || number%2 == 1 {
		fmt.Println(number)
	} else {
		firstNum, SecondNum := 2, number/2
		for {
			// 开始分解自然数
			result = min(result, firstNum+SecondNum)
			if _, ok := mp[SecondNum]; ok && mp[SecondNum] < SecondNum {
				result += mp[SecondNum] - SecondNum
			}
			if SecondNum%2 == 1 {
				break
			}
			firstNum, SecondNum = firstNum*2, SecondNum/2
		}
		fmt.Println(result)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
