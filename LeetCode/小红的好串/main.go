package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str := ""
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str = scanner.Text()
	length := len(str)
	lengthArr := make([]int, 0)
	index := 1
	for index <= length {
		lengthArr = append(lengthArr, index)
		index += 2
	}
	sum := length
	// 滑动窗口
	for i := 1; i < len(lengthArr); i++ {
		tempLength := lengthArr[i]
		left, right := 0, tempLength-1
		//arr := [26]int{}
		mp := map[byte]int{}
		// 一个字符出现次数为奇数, 其他字符出现次数为偶数
		for j := left; j <= right; j++ {
			//arr[str[j]-'a']++
			mp[str[j]]++
		}
		for right < length {
			// 先判断是否满足条件
			if isValid(mp) {
				sum++
			}
			// 窗口右移
			mp[str[left]]--
			if right+1 < length {
				mp[str[right+1]]++
			}
			left++
			right++
		}
	}
	fmt.Println(sum)
}

func isValid(arr map[byte]int) bool {
	flag := false
	for _, value := range arr {
		if value&1 == 1 {
			if flag {
				return false
			}
			flag = true
		}
	}
	return true
}
