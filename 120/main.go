package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	dp := make([][]int, len(triangle)) // dp表示到达第i层的第j个元素的最小路径和
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}
	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		dp[i][0] = dp[i-1][0] + triangle[i][0]
		for j := 1; j < len(triangle[i]); j++ {
			if j > len(dp[i-1])-1 {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}
	minValue := dp[len(dp)-1][0]
	for i := 1; i < len(dp[len(dp)-1]); i++ {
		minValue = min(minValue, dp[len(dp)-1][i])
	}
	return minValue
}

func main() {
	v := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	fmt.Println(v)
}
