package _4__最长特殊序列_II

func findLUSlength(strs []string) int {
	ans := -1
next:
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(strs); j++ {
			if i != j && isSubStr(strs[i], strs[j]) {
				continue next
			}
		}
		if len(strs[i]) > ans {
			ans = len(strs[i])
		}
	}
	return ans
}

func isSubStr(subStr, fullStr string) bool {
	i, j := 0, 0
	for ; i < len(subStr) && j < len(fullStr); j++ {
		if subStr[i] == fullStr[j] {
			i++
		}
	}
	// 最后循环出来 如果 i == len 那么就说明匹配成功
	// 如果 i < len 那么就说明没有匹配成功
	if i == len(subStr) {
		return true
	}
	return false
}
