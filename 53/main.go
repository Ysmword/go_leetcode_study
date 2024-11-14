package main

import (
	"fmt"
	"math"
)

// 贪心算法
func maxSubArray(nums []int) int {
	maxSum := math.MinInt
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if sum < 0 {
			maxSum = max(maxSum, sum)
			sum = 0
		} else {
			maxSum = max(maxSum, sum)
		}
	}
	return maxSum
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
	nums = []int{1}
	fmt.Println(maxSubArray(nums))
	nums = []int{5, 4, -1, 7, 8}
	fmt.Println(maxSubArray(nums))
	nums = []int{-1, -2}
	fmt.Println(maxSubArray(nums))
}
