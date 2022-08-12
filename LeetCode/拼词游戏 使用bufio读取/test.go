package main

import "fmt"

func main() {
	length := 2
	strArr := make([]string, length)
	for i2 := 0; i2 < length; i2++ {
		fmt.Scan(&strArr[i2])
	}
	fmt.Println(strArr)
}
