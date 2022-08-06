package main

import (
	"fmt"
	"math"
)

func main() {
	// 训练集 和 测试集
	pracice := make([]int, 0)
	test := make([]int, 0)
	Bian, LeiBie := 0, 0
	fmt.Scan(&Bian, &LeiBie)
	arr := make([]int, Bian)
	mp := map[int]int{}
	for i := 0; i < Bian; i++ {
		fmt.Scan(&arr[i])
		mp[arr[i]]++
	}
	for i, v := range mp {
		mp[i] = int(math.Ceil(float64(v) / 2.0))
	}
	for i := 0; i < Bian; i++ {
		if mp[arr[i]] > 0 {
			pracice = append(pracice, i+1)
			mp[arr[i]]--
		} else {
			test = append(test, i+1)
		}
	}
	for i := 0; i < len(pracice); i++ {
		fmt.Print(pracice[i], " ")
	}
	fmt.Println()
	for i := 0; i < len(test); i++ {
		fmt.Print(test[i], " ")
	}
}
