package main

import "fmt"

// 接雨水
func trap(height []int) int {
	// 核心思想：对于数组 height 中的每个元素，分别向左和向右扫描并记录左边和右边的最大高度，然后计算每个下标位置能接的雨水量
	leftMax, rightMax := make([]int, len(height)), make([]int, len(height))
	leftMax[0] = height[0]
	rightMax[len(height)-1] = height[len(height)-1]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	sum := 0
	for i := 0; i < len(height); i++ {
		sum += min(leftMax[i], rightMax[i]) - height[i]
	}
	return sum
}

func trap1(heights []int) int {
	left, right := make([]int, len(heights)), make([]int, len(heights))
	left[0] = heights[0]
	for i := 1; i < len(heights); i++ {
		left[i] = max(left[i-1], heights[i])
	}

	right[len(heights)-1] = heights[len(heights)-1]
	for i := len(heights) - 2; i >= 0; i-- {
		right[i] = max(right[i+1], heights[i])
	}

	sum := 0
	for i := 0; i < len(heights); i++ {
		v := min(left[i], right[i]) - heights[i]
		if v > 0 {
			sum = sum + v
		}
	}
	return sum
}

func main() {
	fmt.Println(trap1([]int{4, 2, 0, 3, 2, 5}))
}
