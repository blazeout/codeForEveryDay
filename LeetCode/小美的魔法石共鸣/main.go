package main

import (
	"fmt"
)

/*
  小美有n块魔法石，每块魔法石都有正反两面，每一面上都刻有一个魔法阵，初始状态下，n块魔法石都是正面向上。
  这n块魔法石的能量刚好可以构建一个大型魔法阵，但是需要至少一半的魔法石向上的一面铭刻的阵法相同才能触发大型魔法阵的效果。
  小美希望翻转最少数量的魔法石，使得这个大型魔法阵被触发，请问她最少需要翻转多少块魔法石。
*/

func main() {
	length := 0
	fmt.Scan(&length)
	// 构造魔法阵需要的能量
	sum := 0
	// 正面的魔法石的个数
	positive := make([]int, length)
	// 反面的魔法石的个数
	negative := make([]int, length)
	mp := make(map[int]int)
	for i := 0; i < length; i++ {
		fmt.Scan(&positive[i])
		sum += positive[i]
		mp[positive[i]]++
	}
	for i := 0; i < length; i++ {
		fmt.Scan(&negative[i])
	}
	maxTime := 0
	for _, times := range mp {
		if times > maxTime {
			maxTime = times
		}
	}
	if maxTime > length/2 {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}
}
