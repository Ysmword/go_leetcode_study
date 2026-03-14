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
			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}
		}
	}
	return s[start : start+maxLen]
}

func main() {
	// fmt.Println(longestPalindrome("a"))

	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	// fmt.Println(longestPalindrome(""))
	fmt.Println(longestPalindrome("a"))

	fmt.Println("=====1=====")
	fmt.Println(longestPalindrome1("babad"))
	fmt.Println(longestPalindrome1("cbbd"))
	fmt.Println(longestPalindrome1("a"))
}

func longestPalindrome1(s string) string {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	maxVal := 1
	start := 0
	for j := 0; j < n; j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i <= 2 { // 这里不能限制==2,因为存在bb相连的情况
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i][j] && dp[i+1][j-1]
				}
			}

			if dp[i][j] && maxVal < j-i+1 {
				start = i
				maxVal = j - i + 1
			}
		}
	}
	return s[start : start+maxVal]
}
