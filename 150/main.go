package main

import (
	"fmt"
	"strconv"
)

func calc(a, b string, op string) int {
	v1, _ := strconv.Atoi(a)
	v2, _ := strconv.Atoi(b)
	switch op {
	case "+":
		return v1 + v2
	case "-":
		return v1 - v2
	case "*":
		return v1 * v2
	case "/":
		return v1 / v2
	}
	return 0
}

func isOp(token string) bool {
	if token == "+" || token == "-" || token == "*" || token == "/" {
		return true
	}
	return false
}

// evalRPN 逆波兰表达式求值, 使用栈进行实现
func evalRPN(tokens []string) int {
	stack := []string{}
	for _, token := range tokens {
		if isOp(token) {
			b := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			a := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, strconv.Itoa(calc(a, b, token)))
		} else {
			stack = append(stack, token)
		}
	}
	val, _ := strconv.Atoi(stack[0])
	return val
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
