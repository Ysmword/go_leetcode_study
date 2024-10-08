package main

import "fmt"

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, n+1) // dp表示第i个台阶的方法数
	dp[1] = 1
	dp[2] = 2
	for i := 3; i < len(dp); i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(4))
}
