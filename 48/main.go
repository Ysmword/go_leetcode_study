package main

import "fmt"

// 不按题目要求实现：竖遍历，横赋值，这样无法赋值给matrix
func rotate1(matrix [][]int) {
	res := make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		arr := make([]int, len(matrix))
		for j := 0; j < len(matrix); j++ {
			arr[j] = matrix[j][i]
		}
		res = append(res, arr)
	}
	copy(matrix, res)
}

// 按题目实现：规律：对于矩阵中第 i 行的第 j 个元素，在旋转后，它出现在倒数第 i 列的第 j 个位置。
// matrix[row][col]旋转后变成matrix[col][n−row−1]
func rotate(matrix [][]int) {
	n := len(matrix)
	// 遍历个数，偶数遍历个数：n/2*(n/2)，奇数遍历个数：n/2*((n+1)/2)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			// 四个位置交换，注意顺序，否则会覆盖
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

// 对称翻转，然后在左右翻转
func rotate2(matrix [][]int) {

	// 对称翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 左右翻转
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[i])-j-1] = matrix[i][len(matrix[i])-j-1], matrix[i][j]
		}
	}
}

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	printMatrix(matrix)
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}
}
