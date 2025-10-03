package main

import "fmt"

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	left, right, top, bottom := 0, n-1, 0, n-1
	num := 1
	// 遵守原则：左闭右闭
	for left <= right && top <= bottom {
		// 向右遍历
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++

		// 向下遍历
		for i := top; i <= bottom; i++ {
			res[i][right] = num
			num++
		}
		right--

		// 向左遍历
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--

		// 向上遍历
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

func main() {
	fmt.Println(generateMatrix(3))
}
