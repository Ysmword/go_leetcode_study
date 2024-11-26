package main

import "fmt"

// 简单的位运算就可
func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		if n == n|1 {
			count++
		}
		n = n >> 1
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(2147483645))
}
