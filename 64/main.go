package main

import "fmt"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid)) // 表示到第i行第j列的最小路径和
	for i := 0; i < len(grid); i++ {
		dp[i] = make([]int, len(grid[i]))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			} else if i-1 < 0 && j-1 >= 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if i-1 >= 0 && j-1 < 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			}
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 2, 3},
		{4, 5, 6},
	}))
}
