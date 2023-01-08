package _4_前缀字符串

func prefixCount(words []string, pref string) int {
	ret := 0
	for _, word := range words {
		flag := false
		if len(word) < len(pref) {
			continue
		}
		for i := 0; i < len(pref); i++ {
			if pref[i] != word[i] {
				flag = true
				break
			}
		}
		if flag {
			continue
		}
		ret++
	}
	return ret
}
