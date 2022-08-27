package main

import "fmt"

func main() {
	existMap := map[byte]bool{'A': true, 'H': true, 'I': true, 'M': true, 'O': true, 'T': true, 'U': true,
		'V': true, 'W': true, 'X': true, 'Y': true}
todo:
	for {
		s := ""
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		for i := 0; i < len(s); i++ {
			if !existMap[s[i]] {
				fmt.Println("NO")
				continue todo
			}
		}
		// 比较完就反转
		tempStr := reverse(s)
		if tempStr == s {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
