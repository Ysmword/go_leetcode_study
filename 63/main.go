package main

import "fmt"

// 跟62题一样，只不过多了障碍物，障碍物的路径数为0，其他的状态转移方程不变
// 而且这道题跟青蛙跳台阶一
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp := make([][]int, len(obstacleGrid)) // 表示到第i行第j列的路径数
	for i := 0; i < len(obstacleGrid); i++ {
		dp[i] = make([]int, len(obstacleGrid[i]))
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[i]); j++ {
			if i == 0 && j == 0 && obstacleGrid[i][j] == 0 {
				dp[i][j] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			} else if i-1 < 0 && j-1 >= 0 {
				dp[i][j] = dp[i][j-1]
			} else if i-1 >= 0 && j-1 < 0 {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
}
