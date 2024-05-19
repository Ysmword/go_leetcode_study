package main

import "fmt"

// 方法1：走到最大边界后才跳
func jump(nums []int) int {
	rightMost, end, steps := 0, 0, 0
	for i := 0; i < len(nums)-1; i++ {
		rightMost = max(rightMost, i+nums[i])
		if i == end { // 走到最大边界后才跳
			end = rightMost
			steps++
		}
	}
	return steps
}

func main() {
	fmt.Println(jump([]int{1, 1, 1, 1}))
}


// 2,3,1,1,4
// 第一步 2
// 第二步 3,1
// 第三部 1,4
