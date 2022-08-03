package main

import (
	"fmt"
	"strconv"
)

// 普通方法
func multiply(num1 string, num2 string) (ans string) {
	// 将 70 的 0 直接转移到 curNumber上面去
	m, n := len(num1), len(num2)

	for i := n - 1; i >= 0; i-- {
		curr := ""
		var y, carry int
		y = int(num2[i] - '0')
		for j := m - 1; j >= 0; j-- {
			var x int
			x = int(num1[j] - '0')
			sum := x*y + carry
			curr = strconv.Itoa(sum%10) + curr
			carry = sum / 10
		}
		// 处理进位 carry
		for ; carry != 0; carry /= 10 {
			curr = strconv.Itoa(carry%10) + curr
		}
		tempIndex := i
		// 直接在最后 curr拼接 n个0即可
		for tempIndex < n {
			tempIndex++
			if tempIndex == n {
				break
			}
			curr += "0"
		}
		ans = addStrings(ans, curr)
	}
	return
}

func addStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0
	ans := ""
	for ; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

// 优化竖式相乘

func main() {
	back([]int{1, 2, 3}, 3, []int{})
	fmt.Println(ans)
}

var ans [][]int

func back(data []int, lenth int, res []int) {
	if len(res) == lenth {
		tmp := make([]int, len(res))
		copy(tmp, res)
		ans = append(ans, tmp)
	}

	for i := 0; i < len(data); i++ {
		ans2 := append(res, data[i])
		x := data[:i]
		y := data[i+1:]
		ans1 := append(x, y...)
		back(ans1, lenth, ans2)
	}
}
