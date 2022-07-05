package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) (ans int) {
	existMap := make(map[byte]int, 0)
	left, right := 0, -1
	for ; left < len(s); left++ {
		if left != 0 {
			existMap[s[left-1]]--
		}
		// 一直循环加到出现重复字符
		for right+1 < len(s) && existMap[s[right+1]] == 0 {
			right++
			existMap[s[right]]++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

func lengthOfLongestString(s string) (ans int) {
	positionOfBytes := map[byte]int{}
	left, right := 0, 0
	for right < len(s) {
		if _, ok := positionOfBytes[s[right]]; ok {
			// 如果存在就要更新left的值, 保证更新left 不能小于上次出现的位置 比如abba
			if positionOfBytes[s[right]]+1 > left {
				left = positionOfBytes[s[right]] + 1
			}
			positionOfBytes[s[right]] = right
		} else {
			positionOfBytes[s[right]] = right
		}
		ans = max(ans, right-left+1)
		right++
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	substring := lengthOfLongestSubstring("pwwkew")
	fmt.Println(substring)
}
