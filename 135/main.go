package main

import "fmt"

func candy(ratings []int) int {
	left := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		left[i] = 1
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		}
	}

	right := make([]int, len(ratings))
	for i := len(ratings) - 1; i >= 0; i-- {
		right[i] = 1
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right[i] = right[i+1] + 1
		}
	}

	// 计算最终结果，为什么取最大值呢？因为要保证孩子同时满足左边和右边的情况
	res := 0
	for i := 0; i < len(ratings); i++ {
		res += max(left[i], right[i])
	}
	return res
}

func main() {
	fmt.Println(candy([]int{1, 0, 2}))
}
