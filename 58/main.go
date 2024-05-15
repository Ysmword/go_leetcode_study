package main

import (
	"fmt"
	"strings"
)

// 方法1: 切割，从后面遍历即可
func lengthOfLastWord(s string) int {
	arr := strings.Split(s, " ")
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != "" {
			return len(arr[i])
		}
	}
	return 0
}

// 方法2:逆向思维，从后面开始遍历，直到没有遇到空，如果遇到空就开始记录，如果再次遇到空
// 直接返回
func lengthOfLastWord1(s string) int {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}

	ans := 0
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLastWord1("   fly me   to   the moon  "))
}
