package main

import "fmt"

func trailingZeroes(n int) int {
	if n == 0 {
		return 0
	}
	// 同化为：n/5 + n/5/5 + n/5/5/5 + n/5/5/5/5 + n/5/5/5/5/5 + n/5/5/5/5/5/5 +...
	return n/5 + trailingZeroes(n/5) // 10 = 2 * 5, 因此0的个数等于n/5
}

func main() {
	fmt.Println(trailingZeroes(5))
}
