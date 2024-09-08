package main

import "fmt"

// 回溯算法，超时严重，哈哈哈
func lengthOfLIS(nums []int) int {
	choic := make([]int, 0)
	maxLen := 0
	backtrack(0, &maxLen, nums, choic)
	return maxLen
}

func backtrack(start int, maxLen *int, nums []int, choic []int) {
	// if start == len(nums) { 不能在这里判断，因为有一些情况可能都不会到达这里
	// 	*maxLen = max(*maxLen, len(choic))
	// 	return
	// }

	for i := start; i < len(nums); i++ {
		if len(choic) == 0 || nums[i] > choic[len(choic)-1] {
			choic = append(choic, nums[i])
			backtrack(i+1, maxLen, nums, choic)
			choic = choic[:len(choic)-1]
		}
	}
	*maxLen = max(*maxLen, len(choic))
}

// 动态规划，反过来考虑，必须以i为结尾，dp[i] = max(dp[i],max(dp[i-1]+1,dp[i-2]+1...dp[0]+1))
func lengthOfLIS1(nums []int) int {
	dp := make([]int, len(nums))
	var maxAns int = 1
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxAns = max(maxAns, dp[i])
	}
	return maxAns
}

func main() {
	// fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	// fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	// fmt.Println(lengthOfLIS([]int{7, 7, 7, 7}))
	// fmt.Println(lengthOfLIS([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))
	fmt.Println(lengthOfLIS1([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS1([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS1([]int{7, 7, 7, 7}))
	fmt.Println(lengthOfLIS1([]int{1, 3, 6, 7, 9, 4, 10, 5, 6}))

}
