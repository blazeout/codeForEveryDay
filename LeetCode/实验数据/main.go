package main

import (
	"strconv"
)

func moveZeros(nums []int) {
	left, right, length := 0, 0, len(nums)
	for right < length {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeros(nums)
	println(nums)
}

func fixIndex(b []byte, i int) {
	if b[i] <= '9' && b[i] >= '0' {
		return
	}
	var num1, num2 int
	num1, num2 = int(b[i-1]-'0'), int(b[i+1]-'0')
	for j := 0; j < 10; j++ {
		if j != num1 && j != num2 {
			b[i] = byte(j) + '0'
			return
		}
	}
	return
}

func fixIndexLast(b []byte, i int) {
	if b[i] <= '9' && b[i] >= '0' {
		return
	}
	num := int(b[i-1] - '0')
	for j := 0; j < 10; j++ {
		if j != num {
			b[i] = byte(j) + '0'
			atoi, _ := strconv.Atoi(string(b))
			if atoi%3 == 0 {
				return
			} else {
				b[i] = '?'
				continue
			}
		}
	}
}

func fixIndexOne(b []byte, i int) {
	// 不能是前导0
	if b[i] <= '9' && b[i] >= '1' {
		return
	}
	if b[i] == '?' && b[i+1] != '1' {
		b[i] = '1'
		return
	}
	if b[i] == '?' && b[i+1] == '1' {
		b[i] = '2'
	}
	return
}
