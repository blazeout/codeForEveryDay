package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		// 如果是数字 则直接入栈
		atoi, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, atoi)
		} else {
			// 如果是运算符 则取出两个数字 做运算 再入栈
			var a, b int
			if len(stack) >= 2 {
				a = stack[len(stack)-2]
				b = stack[len(stack)-1]
				stack = stack[:len(stack)-2]
			}
			switch tokens[i] {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				stack = append(stack, a/b)
			}
		}
	}
	return stack[0]
}

func main() {
	rpn := evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"})
	fmt.Println(rpn)
}
