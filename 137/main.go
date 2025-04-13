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

func singleNumber1(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ {
		var total int32
		for j := 0; j < len(nums); j++ {
			bit := int32(nums[j]) >> i & 1
			total += bit
		}
		if total%3 > 0 {
			res = res | 1<<i
		}
	}
	return int(res)
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber1([]int{2, 2, 3, 2}))
}
