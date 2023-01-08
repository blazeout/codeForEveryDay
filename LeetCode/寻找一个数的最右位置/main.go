package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param nums int整型一维数组
 * @param target int整型
 * @return int整型
 */
func searchRight(nums []int, target int) int {
	// write code here
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if right < 0 || nums[right] != target {
		return -1
	}
	return right
}

func main() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		println(err)
	//	}
	//	fmt.Println("end")
	//}()
	//
	//go func(i int) {
	//	defer func() {
	//		if err := recover(); err != nil {
	//			println("go func", err)
	//		}
	//		fmt.Println("go func panic catched")
	//	}()
	//	if i > 1 {
	//		panic("i > 1")
	//	}
	//}(1)
	//
	//ch := make(chan int, 2)
	//fmt.Println(cap(ch))
	//arr := [1]int{1}
	//fmt.Println(cap(arr))
	//slice := &[]int{1}
	//fmt.Println(cap(*slice))

	f()
	fmt.Println("O")
}

func f() {
	defer func() {
		fmt.Println("M")
	}()
	fmt.Println("N")
}
