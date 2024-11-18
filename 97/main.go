package main

import "fmt"

// 最简单的做法：统计判断，不行，还得看顺序
func isInterleaveFail(s1 string, s2 string, s3 string) bool {
	dp1, dp2, dp3 := make(map[byte]int), make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		dp1[s1[i]]++
	}
	for i := 0; i < len(s2); i++ {
		dp2[s2[i]]++
	}
	for i := 0; i < len(s3); i++ {
		dp3[s3[i]]++
	}
	for k, v := range dp3 {
		if dp1[k]+dp2[k] != v {
			return false
		}
	}
	return false
}

// 三指针，不行，因为取值的顺序有问题
func isInterleaveFail1(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	x, y, k := 0, 0, 0
	for k < len(s3) {
		if x < len(s1) && s1[x] == s3[k] {
			x++
			k++
		} else if y < len(s2) && s2[y] == s3[k] {
			y++
			k++
		} else {
			return false
		}
	}
	return true
}

// 动态规划
// 方程：dp[i][j] = ((s1[i]==s3[i+j])&&dp[i-1][j]) || ((s2[j]==s3[i+j])&&dp[i][j-1])
// dp[i][j]的定义为：表示 s1的前i个元素和s2的前j个元素是否能交错组成s3的前 i+j 个元素。
// 初始化：dp[0][0] = true，因为两个空字符串可以交错组成一个空字符串。

func isInterleave(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if n+m != t {
		return false
	}
	dp := make([][]bool, n+1) // 为啥n+1,m+1,因为是代表为个
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	dp[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return dp[n][m]
}

func main() {
	fmt.Println(isInterleaveFail("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleaveFail1("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
}
