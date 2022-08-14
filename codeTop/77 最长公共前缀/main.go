package main

func longestCommonPrefix(strs []string) string {
	// 纵向比较
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			// 当 i 等于某一个字符串的长度时 就可以直接返回了
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	// 上面的 for 循环完了 那么就说明str[0] 里面的字符都在 就可以返回了
	return strs[0]
}

func main() {

}
