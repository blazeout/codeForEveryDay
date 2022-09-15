package main

import "fmt"

func main() {
	//length := 0
	//fmt.Scan(&length)
	//mp := map[byte]int{}
	//mp['a'] = 1
	//for i := 'b'; i <= 'z'; i++ {
	//	mp[byte(i)] = 2 * mp[byte(i-1)]
	//}
	//str := ""
	//for length >= 1 {
	//	for i := 'z'; i >= 'a'; i-- {
	//		if length >= mp[byte(i)] {
	//			str += string(i)
	//			length -= mp[byte(i)]
	//			continue
	//		}
	//	}
	//}
	//fmt.Println(str)

	arr := [26]int{}
	arr[0] = 1
	arr[1] = 2
	flag := false
	for index, value := range arr {
		fmt.Println(string(rune(index+'a')), value)
		if value&1 == 1 {
			if flag {
				fmt.Println("false")
				return
			}
			flag = true
		}
	}
}
