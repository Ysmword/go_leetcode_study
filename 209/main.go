package main

import (
	"fmt"
	"math"
)

// 方法1：滑动窗口，双指针，随着双指针所有数之和，移动左右指针，如果大于target，移动左指针，反之移动右指针
func minSubArrayLen(target int, nums []int) int {
	start, end := 0, 0
	sum := 0
	ans := math.MaxInt
	for end < len(nums) {
		sum += nums[end]
		for sum >= target {
			ans = min(ans, end-start+1)
			sum -= nums[start]
			start++
		}
		end++
	}
	if ans == math.MaxInt {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))
}
