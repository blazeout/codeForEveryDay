package main

import "fmt"

// 小美在做一个实验，这个实验会在纸带上打印出若干个数字，已知该实验所呈现出的正确结果应该是存在某一个分割处k，在k之前打印出的数字都是小于0的，而在k之后的数字应该都是大于0的，那么在k之前如果某一个数据大于等于0，那么我们认为这个数据是异常的，同理，在k之后如果某一个数据小于等于0，那么我们也认为这个数据是异常的。
// 现在给出小美打印的纸带，且k是未知的，那么请问在最乐观的情况下至少有多少个实验数据是异常的。(显然如果出现0，无论k为哪个时刻，这个0数据都是异常的
func main() {
	length := 0
	fmt.Scan(&length)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		fmt.Scan(&arr[i])
	}
	// 滑动窗口 首先计算异常个数
	count1, count2 := 0, 0
	if arr[0] >= 0 {
		count1++
	}
	for i := 1; i < length; i++ {
		if arr[i] <= 0 {
			count2++
		}
	}
	sum := count1 + count2
	index := 0
	for i := 1; i < length; i++ {
		if arr[i] >= 0 {
			count1++
		}
		if arr[i] <= 0 {
			count2--
		}
		if count1+count2 < sum {
			sum = count1 + count2
			index = i
		}
	}
	fmt.Println(index + 1)
}
