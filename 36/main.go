package main

import "fmt"

// 根据题意可以分成3部分进行校验
// 按行校验，按列校验，按3x3校验
// 1. 按行校验 （i,j）i表示行，j表示列
// 1.1 行内保证无重复数字，row[i][val]<=1
// 2. 按列校验
// 2.1 列内保证无重复数字，col[j][val]<=1
// 3. 按3x3校验
// // 3.1 3x3内保证无重复数字，box[i/3][j/3][val]<=1
func isValidSudoku(board [][]byte) bool {
	row := make([][]int, 9)
	col := make([][]int, 9)
	box := make([][][]int, 3)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				if len(row[i]) == 0 {
					row[i] = make([]int, 10)
				}
				if row[i][board[i][j]-48] > 0 {
					return false
				}
				row[i][board[i][j]-48]++

				if len(col[j]) == 0 {
					col[j] = make([]int, 10)
				}
				if col[j][board[i][j]-48] > 0 {
					return false
				}
				col[j][board[i][j]-48]++

				if len(box[i/3]) == 0 {
					box[i/3] = make([][]int, 3)
				}
				if len(box[i/3][j/3]) == 0 {
					box[i/3][j/3] = make([]int, 10)
				}
				if box[i/3][j/3][board[i][j]-48] > 0 {
					return false
				}
				box[i/3][j/3][board[i][j]-48]++
			}
		}
	}
	return true
}

func main() {
	a := [][]byte{{'5', '3', '8'}, {'.', '.', '.'}, {'8', '.', '.'}}
	fmt.Println(isValidSudoku(a))
}

// 矩阵转化为一位数组
