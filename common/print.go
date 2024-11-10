package common

import "fmt"

// PrintMatrix 打印矩阵
func PrintMatrix[T int | int64 | int32 | float32 | float64](matrix [][]T) {
	for row := 0; row < len(matrix); row++ {
		fmt.Println(matrix[row])
	}
}
