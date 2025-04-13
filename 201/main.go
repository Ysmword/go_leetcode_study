package main

import (
	"fmt"
)

// 有意思，这样其实就会变成常数遍历
func rangeBitwiseAnd(left int, right int) int {
	for left < right {
		right = right & (right - 1)
	}
	return right
}

// 找公共前缀
func rangeBitwiseAnd1(m int, n int) int {
	shift := 0
	for m < n {
		m, n = m>>1, n>>1
		shift++
	}
	return m << shift
}

func main() {
	fmt.Println(rangeBitwiseAnd(1, 9))
	fmt.Println(rangeBitwiseAnd1(1, 9))
}
