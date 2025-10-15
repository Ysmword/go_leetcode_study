package main

import "fmt"

func largestRectangleArea(heights []int) int {
	res := 0
	heights = append(heights, 0)           // 追加一个0，防止出现从大到小的情况
	heights = append([]int{0}, heights...) // 追加一个0，防止出现从小到大的情况
	stack := make([]int, 0)                // 单调栈，存放高度的下标，单调递增

	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] { // 出现小于栈顶的情况
			// 从栈顶开始，计算以栈顶为高度的矩形面积
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			right := i
			width := right - left - 1
			res = max(res, height*width)
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	res := largestRectangleArea(heights)
	fmt.Println(res)
}
