package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	length1 := scanner.Text()
	length, _ := strconv.Atoi(length1)
	scanner.Scan()
	splitArr := strings.Split(scanner.Text(), " ")
	intArr := make([]int, len(splitArr))
	for i := 0; i < len(splitArr); i++ {
		atoi, _ := strconv.Atoi(splitArr[i])
		intArr[i] = atoi
	}
	scanner.Scan()
	compareTime1 := scanner.Text()
	compareTime, _ := strconv.Atoi(compareTime1)
	scanner.Scan()
	compareIndexArr1 := strings.Split(scanner.Text(), " ")
	compareIndexArr := make([]int, len(compareIndexArr1))
	for i := 0; i < len(compareIndexArr); i++ {
		temp, _ := strconv.Atoi(compareIndexArr1[i])
		compareIndexArr[i] = temp
	}
	// 开始比较
	for i := 0; i < compareTime; i++ {
		temp1 := []int{}
		temp2 := []int{}
		index := compareIndexArr[i]
		for j := 0; j < length; j++ {
			if intArr[j] < index {
				temp1 = append(temp1, intArr[j])
			} else if intArr[j] > index {
				temp2 = append(temp2, intArr[j])
			}
		}
		temp1 = append(temp1, index)
		temp1 = append(temp1, temp2...)
		intArr = temp1
	}
	for i := 0; i < len(intArr); i++ {
		if i != length-1 {
			fmt.Print(intArr[i], " ")
		} else {
			fmt.Print(intArr[i])
		}
	}
}
