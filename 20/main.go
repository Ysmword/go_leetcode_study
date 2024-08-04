package main

import "fmt"

// 栈
func isValid(s string) bool {
	stack := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			stack = append(stack, string(s[i]))
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if string(s[i]) == ")" && stack[len(stack)-1] != "(" {
			return false
		}
		if string(s[i]) == "}" && stack[len(stack)-1] != "{" {
			return false
		}
		if string(s[i]) == "]" && stack[len(stack)-1] != "[" {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return true && len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(}"))
}
