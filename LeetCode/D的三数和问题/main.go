package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	times := stringToInt(scanner.Text())
	arrStr := make([]string, times)
	for i := 0; i < times; i++ {
		length := 0
		scanner.Scan()
		length = stringToInt(scanner.Text())
		arr := make([]int, length)
		scanner.Scan()
		split := strings.Split(scanner.Text(), " ")
		for j := 0; j < length; j++ {
			arr[j] = stringToInt(split[j])
		}
		res := isSumOfThree(arr)
		arrStr[i] = res
	}
	for i := 0; i < len(arrStr); i++ {
		fmt.Println(arrStr[i])
	}
}

func isSumOfThree(arr []int) string {
	mp := map[int]int{}
	for i := 0; i < len(arr); i++ {
		mp[arr[i]] = i
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for z := j + 1; z < len(arr); z++ {
				if _, ok := mp[arr[i]+arr[j]+arr[z]]; !ok && mp[arr[i]+arr[j]+arr[z]] < 1 {
					return "NO"
				}
			}
		}
	}
	return "YES"
}
