package main

import "fmt"

// 循环+哈希表，相遇原理：数字不会无限大，只会在有限的集合中循环
func isHappy(n int) bool {
	exist := make(map[int]bool)
	for !exist[n] {
		exist[n] = true
		n = valSplit(n)
		if n == 1 {
			return true
		}
	}
	return false
}

func valSplit(n int) int {
	val := 0
	for n >= 10 {
		re := n % 10
		val += re * re
		n = n / 10
	}
	return val + n*n
}

// 快慢指针,极致
func isHappy2(n int) bool {
	s, f := n, valSplit(n)
	for s != f && f != 1 {
		s = valSplit(s)
		f = valSplit(valSplit(f))
	}
	return s == 1
}

func main() {
	fmt.Println(isHappy(2))
	fmt.Println(valSplit(100))
	fmt.Println(isHappy2(2))
	fmt.Println(valSplit(100))
}
