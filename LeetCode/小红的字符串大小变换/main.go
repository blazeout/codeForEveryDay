package main

import (
	"fmt"
	"strings"
)

// 京东第一题 8.27
func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	str := ""
	fmt.Scan(&str)
	upper := strings.ToUpper(str[:k])
	lower := strings.ToLower(str[k:])
	fmt.Println(upper + lower)
}
