package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[1], nums[0])
	}
	dp := make([]int, len(nums)+1) // dp表示第i家的最大金额
	dp[1] = nums[0]
	dp[2] = max(nums[1], nums[0])
	for i := 2; i < len(nums); i++ {
		dp[i+1] = max(dp[i-1]+nums[i], dp[i])
	}
	return dp[len(nums)]
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
