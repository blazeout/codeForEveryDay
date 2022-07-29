package _5_有效的正方形

import "sort"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	// 计算 边和 对角线的长度
	cache := []int{calc(p1, p2), calc(p2, p3), calc(p3, p4), calc(p1, p4), calc(p1, p3), calc(p2, p4)}
	sort.Ints(cache)
	return cache[0] > 0 && cache[0] == cache[1] && cache[1] == cache[2] && cache[2] == cache[3] && cache[4] == cache[5]
}

func calc(p1, p2 []int) int {
	return (p2[0]-p1[0])*(p2[0]-p1[0]) + (p2[1]-p1[1])*(p2[1]-p1[1])
}
