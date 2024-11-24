package main

import (
	"fmt"
	"math"
)

// 找排序
func findMin(nums []int) int {
	minVal := math.MaxInt
	x, y := 0, len(nums)-1
	for x <= y {
		mid := (x + y) / 2
		if nums[x] <= nums[mid] { // 说明[0,mid]是有序的
			minVal = min(nums[x], minVal)
			x = mid + 1
		} else { // 说明[mid,len(nums)-1]是有序的
			minVal = min(nums[mid], minVal)
			y = mid - 1
		}
	}
	return minVal
}

func main() {
	fmt.Println(findMin([]int{2, 1}))
}
