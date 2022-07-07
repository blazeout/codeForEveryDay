package main

// 实现一个函数，检查一个字符串是否是另一个字符串的子序列。返回第一次出现的位置，如果没有返回-1。

func strStr(haystack string, needle string) int {
	// write code here
next:
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				continue next
			}
		}
		return i
	}
	return -1
}

func main() {

}
