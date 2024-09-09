package main

import (
	"fmt"
)

// 正常思维会超时
func myPow(x float64, n int) float64 {
	if x == 0 || n == 0 {
		return 1
	}
	temp := x
	if n < 0 {
		for i := -1; i >= n; i = i - 1 {
			if i == -1 {
				x = 1 / temp
			} else {
				x = x * (1 / temp)
			}
		}
		return x
	}
	for i := 2; i <= n; i++ {
		x = x * temp
	}
	return x
}

// 特殊思维：当n为偶数：x^n = y^2，当n为奇数：x^n = y^2 * x
func myPow1(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	fmt.Println(myPow(2, -3))
	fmt.Println(myPow(2.10000, 3))
}
