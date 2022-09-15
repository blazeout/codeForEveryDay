package main

import "fmt"

func main() {
	str := ""
	count := 0
	fmt.Scan(&str)
	mp := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	cnt := [26]int{}
	// 1, 4 字符是辅音 2 3 5 字符是元音
todo:
	for i := 0; i < len(str)-4; i++ {
		for j := 0; j+i < len(str); j++ {
			if j == 5 {
				break
			}
			switch j {
			case 0, 3:
				if _, ok := mp[str[i+j]]; !ok && cnt[str[i+j]-'a'] == 0 {
					cnt[str[i+j]-'a']++
				} else {
					continue todo
				}
			case 1, 2, 4:
				if _, ok := mp[str[i+j]]; ok && cnt[str[i+j]-'a'] == 0 {
					cnt[str[i+j]-'a']++
				} else {
					continue todo
				}
			}

		}
		cnt = [26]int{}
		count++
	}
	fmt.Println(count)
}
