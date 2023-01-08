package main

import "fmt"

// 从键盘输入不确定个数的整数, 以-1结束, 统计其中能被3整除的个数, 能被5整除的个数并输出
// 例如: 输入: 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 -1
// 输出: 能被3整除的个数: 5
// 输出: 能被5整除的个数: 3
func main() {
	var num int
	var count3, count5 int
	for {
		fmt.Scan(&num)
		if num == -1 {
			break
		}
		if num%3 == 0 {
			count3++
		}
		if num%5 == 0 {
			count5++
		}
	}
	fmt.Printf("能被3整除的个数: %d", count3)
	fmt.Printf("能被5整除的个数: %d", count5)
}
