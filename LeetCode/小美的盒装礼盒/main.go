package main

import "fmt"

// 小美开的西点屋举办一周年活动，她准备制作一批礼盒作为对消费者的回馈，每个礼盒中都有三枚西点屋的招牌点心。
// 西点屋共有A和B两种招牌点心，为了让消费者都能品尝到两种点心，
// 因此每个礼盒中都要包含至少一枚A点心和一枚B点心。现在小美的西点屋内共有x枚A点心和y枚B点心，请问小美最多可以制作多少个礼盒。
func calc(x, y int) int {
	if x > y {
		x, y = y, x
	}
	if x*2 <= y {
		return x
	}
	// 大的多组成 2
	temp := y / 2
	y = y - temp*2
	x = x - temp
	for x >= 2 {
		if y >= 1 {
			y--
			temp++
		}
		x -= 2
	}
	return temp
}

func main() {
	T := 0
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		x, y := 0, 0
		fmt.Scan(&x, &y)
		if x <= y/2 {
			fmt.Println(x)
		} else {
			fmt.Println(y/2 + 1)
		}
	}
}
