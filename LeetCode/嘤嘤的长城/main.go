package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	defer func() {
		recover()
	}()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	atoi, _ := strconv.Atoi(text)
	length := atoi
	scanner.Scan()
	split := strings.Split(scanner.Text(), " ")
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i], _ = strconv.Atoi(split[i])
	}
	// 贪心的认为 找出数字最多的和第二多的两个数字, 然后以两个数字为奇偶分别遍历
	mp := map[int]int{}
	for i := 0; i < length; i++ {
		mp[arr[i]]++
	}
	max1, max2 := 0, 0
	element1, element2 := 0, 0
	for i, v := range mp {
		if v > max1 {
			max2 = max1
			element2 = element1
			max1 = v
			element1 = i
		} else if v > max2 {
			max2 = v
			element2 = i
		}
	}
	// 然后以两个数字为奇偶分别遍历
	count, count1 := 0, 0
	for i := 0; i < length; i++ {
		if i%2 == 0 && arr[i] != element1 {
			count++
		} else if i%2 == 1 && arr[i] != element2 {
			count++
		}
		if i%2 == 0 && arr[i] != element2 {
			count1++
		} else if i%2 == 1 && arr[i] != element1 {
			count1++
		}
	}
	fmt.Println(min(count, count1))
}

func min(count int, count1 int) int {
	if count < count1 {
		return count
	}
	return count1
}
