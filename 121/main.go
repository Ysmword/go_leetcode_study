package main

import (
	"fmt"
	"math"
)

// 方法1：暴力破解，循环两次, 会超时
func maxProfit(prices []int) int {
	max := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if max < prices[j]-prices[i] {
				max = prices[j] - prices[i]
			}
		}
	}
	return max
}

// 方法2：贪心算法：在遍历过程中，如果我遇到更低的，我会改变策略，我重新选择买入时间，每次遍历我都记录一次我的收入
func maxProfit1(prices []int) int { // 有意思这个
	max := 0
	min := math.MaxInt // 直接等第一个值就可以
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		if prices[i] > min {
			yingli := prices[i] - min
			if yingli > max {
				max = yingli
			}
		}
	}
	return max
}

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1(nums))
}
