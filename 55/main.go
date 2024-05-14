package main

import "fmt"

/*
1. 如果某一个作为 起跳点 的格子可以跳跃的距离是 3，那么表示后面 3 个格子都可以作为 起跳点
2. 可以对每一个能作为 起跳点 的格子都尝试跳一次，把 能跳到最远的距离 不断更新
3. 如果能跳到最远的距离大于数组长度就一定能跳到
*/

func canJump(nums []int) bool {
	rightMost := 0
	for i, num := range nums {
		if i > rightMost {
			return false
		}
		rightMost = max(rightMost, i+num)
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}
