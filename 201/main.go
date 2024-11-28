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

func rangeBitwiseAnd1(left, right int) int {
	shift := 0
	for left < right {
		left = left >> 1
		right = right >> 1
		shift++
	}
	return right << shift
}

func main() {
	fmt.Println(rangeBitwiseAnd(1, 2147483647))
}
