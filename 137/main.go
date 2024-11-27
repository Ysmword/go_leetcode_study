package main

import "fmt"

// 可以参考二进制：
func singleNumber(nums []int) int {
	dp := make(map[int]int)
	for _, num := range nums {
		dp[num]++
	}
	for key, value := range dp {
		if value == 1 {
			return key
		}
	}
	return -1
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
}
