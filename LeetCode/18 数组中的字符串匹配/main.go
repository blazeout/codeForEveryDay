package _8_数组中的字符串匹配

import "strings"

func stringMatching(words []string) []string {
	arr := make([]string, 0)
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			if strings.Contains(words[j], words[i]) {
				arr = append(arr, words[i])
				break
			}
		}
	}
	return arr
}
