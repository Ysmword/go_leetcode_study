package main

import (
	"fmt"
	"sort"
)

// 排序+子序列循环(双指针)
func longestConsecutive(nums []int) int {
	sort.Ints(nums)
	lg := 0
	for i := 0; i < len(nums); i++ {
		count := 1
		for i < len(nums)-1 {
			if nums[i+1]-nums[i] == 1 {
				count++
				i++
			} else if nums[i+1]-nums[i] == 0 {
				i++
			} else {
				break
			}
		}
		lg = max(lg, count)
	}
	return lg
}

func main() {
	fmt.Println(longestConsecutive([]int{1, 2, 3, 4, 6, 8, 9}))
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	fmt.Println(longestConsecutive([]int{0}))
	fmt.Println(longestConsecutive([]int{1, 2, 0, 1}))
}
