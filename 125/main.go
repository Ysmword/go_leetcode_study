package main

import (
	"fmt"
	"strings"
)

// 先转化为小写，然后过滤非字母数字字符，最后判断是否为回文
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	dest := []byte{}
	for _, item := range s {
		if item >= 48 && item <= 57 {
			dest = append(dest, byte(item))
		}
		if item >= 97 && item <= 122 {
			dest = append(dest, byte(item))
		}
	}
	x, y := 0, len(dest)-1
	for x < y {
		if dest[x] != dest[y] {
			return false
		}
		x++
		y--
	}
	return true
}

// 进阶做法：一次遍历，遇到非字母数字字符，自动过滤

func main() {
	fmt.Println(isPalindrome("0P"))
}
