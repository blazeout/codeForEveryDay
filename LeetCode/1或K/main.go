package main

import "fmt"

//有一个变量a,其初始值为0, 还有一个计数器k,其初始值为1,每次操作, 可以让a 加上计数器k的值或者让a加一
//每次操作之后,计数器k加一,米小优想知道他最少操作多少次可以让a的值等于n`
// 思路: 贪心
func main() {
	t := 0
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		n := 0
		fmt.Scan(&n)
		times := minOperations(n)
		if times > n/2+1 {
			fmt.Println(n/2 + 1)
		} else {
			fmt.Println(times)
		}
	}
}

func minOperations(n int) int {
	// 模拟贪心算法, 每次操作都是加上计数器的值
	// 直到a > n, 就撤销上一次的操作, 并且计数器减一, 换成+1操作
	count := 1
	a := 0
	for a < n {
		a += count
		count++
		if a == n {
			return count - 1
		}
		for a > n {
			// 撤销上一次的操作
			a -= count - 1
			count--
			for a < n {
				a += 1
				count++
			}
		}
		if a == n {
			return count - 1
		}
	}
	return -1
}
