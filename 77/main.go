package main

import "fmt"

// 组合：回溯
func combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	backtrack(1, n, k, []int{}, &ans)
	return ans
}

func backtrack(start, n, k int, temp []int, ans *[][]int) {
	if len(temp) == k { // 是否为解
		// 拷贝一份
		temp2 := make([]int, len(temp))
		copy(temp2, temp)
		*ans = append(*ans, temp2)
		return
	}
	for i := start; i <= n; i++ {
		temp = append(temp, i)          // 做出选择
		backtrack(i+1, n, k, temp, ans) // 进行下一轮选择
		temp = temp[:len(temp)-1]       // 回退：撤销选择，恢复到之前的状态
	}
}

func main() {
	fmt.Println(combine(4, 2))
}
