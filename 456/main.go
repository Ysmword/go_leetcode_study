package main

import (
	"fmt"
	"math"
)

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// 维护一个k的单调递减栈（栈的是j,弹出的才是k），这样能够找到j和i
	stk := make([]int, 0)
	k := math.MinInt // 最小值
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < k {
			return true
		}

		for len(stk) > 0 && stk[len(stk)-1] < nums[i] { // 大于栈顶，出栈
			k = stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, nums[i])
	}
	return false
}

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))  // false
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))  // true
	fmt.Println(find132pattern([]int{-1, 3, 2, 0})) // true
}
