package main

import "fmt"

// dp[i,j]为i种硬币可以凑成总金额所需的最小的硬币个数
func coinChange(coins []int, amount int) int {
	dp := make([][]int, len(coins)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, amount+1)
	}

	for i := 1; i <= amount; i++ {
		dp[0][i] = amount + 1 // 特殊值表示没办法凑出
	}

	for i := 1; i <= len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j < coins[i-1] {
				dp[i][j] = dp[i-1][j] // 不选当前硬币，因为当前硬币，已经超过总金额
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			}
		}
	}
	if dp[len(coins)][amount] > amount {
		return -1
	}
	return dp[len(coins)][amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{186, 419, 83, 408}, 6249))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
}
