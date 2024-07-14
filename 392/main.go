package main

import "fmt"

// 方法1：双指针
func isSubsequence(s string, t string) bool {
	x, y := 0, 0
	for x < len(s) && y < len(t) {
		if s[x] == t[y] {
			x++
		}
		y++
	}

	return len(s) == x
}

func main() {
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
