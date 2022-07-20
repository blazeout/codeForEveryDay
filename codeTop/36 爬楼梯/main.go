package _6_爬楼梯

func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	// a是前置 b是爬1阶共有一种方法, c是爬2阶也有两种方法
	a, b, c := 1, 1, 2
	// 从第三阶开始爬咯
	for i := 3; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}
	return c
}
