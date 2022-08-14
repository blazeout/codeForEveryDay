package main

import (
	"fmt"
)

func main() {
	a, b := "", ""
	fmt.Scan(&a)
	fmt.Scan(&b)
	if a == b {
		fmt.Println(0)
	} else if a < b {
		fmt.Println(-1)
	} else if a > b {
		fmt.Println(1)
	}
	fmt.Printf("%d %d", 'A', 'a')
}
