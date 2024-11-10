package main

import "leetcode/common"

// 遍历收集，然后赋值
func setZeroes(matrix [][]int) {
	rows := make(map[int]struct{})
	cols := make(map[int]struct{})

	// 收集所有要赋值为0的行和列
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			if matrix[row][col] == 0 {
				rows[row] = struct{}{}
				cols[col] = struct{}{}
			}
		}
	}

	for row := range rows {
		for col := 0; col < len(matrix[row]); col++ {
			matrix[row][col] = 0
		}
	}

	for col := range cols {
		for row := 0; row < len(matrix); row++ {
			matrix[row][col] = 0
		}
	}
}

// 方法优化：使用第一行和第一列代替上面的rows和cols

func main() {
	matrix := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(matrix)
	common.PrintMatrix(matrix)
}
