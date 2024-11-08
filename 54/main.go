package main

import "fmt"

// 老实点，右左上下，遍历一圈
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	for len(res) <= len(matrix)*len(matrix[0]) && left <= right && top <= bottom {
		// 向右遍历
		for i := left; i <= right && top <= bottom; i++ {
			fmt.Println("右", matrix[top][i])
			res = append(res, matrix[top][i])
		}

		// 向下遍历
		top++
		for i := top; i <= bottom && left <= right; i++ {
			fmt.Println("下", matrix[i][right])
			res = append(res, matrix[i][right])
		}

		// 向左遍历
		right--
		for i := right; i >= left && top <= bottom; i-- {
			fmt.Println("左", matrix[bottom][i])
			res = append(res, matrix[bottom][i])
		}

		// 向上遍历
		bottom--
		for i := bottom; i >= top && left <= right; i-- {
			fmt.Println("上", matrix[i][left])
			res = append(res, matrix[i][left])
		}
		left++
	}
	return res
}

func main() {
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// res := spiralOrder(matrix)
	// fmt.Println(res)

	matrix2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	res := spiralOrder(matrix2)
	fmt.Println(res)
}
