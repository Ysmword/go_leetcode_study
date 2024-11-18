package main

import "fmt"

// 动态规划
// 方程：dp[i][j] = (s[i]==s[j])&&(j-i<3||dp[i+1][j-1])
// 为什么考虑j-i<3？ 因为如果s[i]==s[j]，且j-i<3，那么s[i+1][j-1]肯定是回文串
// 举个例子：s="aba"，s[0][2]肯定是回文串，因为s[0]==s[2]，且j-i<3
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	// 初始化
	maxLen := 1
	start := 0
	// 遍历列开始
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ { // 为啥是i<j,因为我们dp[i][j]定义
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			// 如果dp[i][j]是回文串，且长度大于maxLen，那么更新maxLen和start
			for dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}

func main() {
	fmt.Println(longestPalindrome("a"))
}
