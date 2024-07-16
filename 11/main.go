package main

import "fmt"

// 双指针，宽固定
func maxArea(height []int) int {
	x, y, maxArea := 0, len(height)-1, 0
	for x < y {
		if height[x] < height[y] {
			// 相当于以height[x]为高，最大面积为(y - x) * height[x]
			maxArea = max((y-x)*height[x], maxArea)
			x++
		} else {
			// 相当于以height[y]为高，最大值为(y - x) * height[y]
			maxArea = max((y-x)*height[y], maxArea)
			y--
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
