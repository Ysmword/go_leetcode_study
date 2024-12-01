package main

import "fmt"

// 回溯算法
func exist(board [][]byte, word string) bool {
	rows, cols, index := len(board), len(board[0]), 0
	visited := make([][]int, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if board[i][j] == word[index] {
				visited[i][j] = 1
				if backtrace(board, word, cols, rows, index+1, j, i, visited) {
					return true
				}
				visited[i][j] = 0
			}
		}
	}
	return false
}

func backtrace(board [][]byte, word string, cols, rows, index int, col, row int, visited [][]int) bool {
	if index == len(word) {
		return true
	}

	// 尝试向右走
	if col+1 < cols && visited[row][col+1] != 1 && board[row][col+1] == word[index] {
		visited[row][col+1] = 1
		if backtrace(board, word, cols, rows, index+1, col+1, row, visited) {
			return true
		}
		visited[row][col+1] = 0
	}

	// 尝试向左走
	if col-1 >= 0 && visited[row][col-1] != 1 && board[row][col-1] == word[index] {
		visited[row][col-1] = 1
		if backtrace(board, word, cols, rows, index+1, col-1, row, visited) {
			return true
		}
		visited[row][col-1] = 0
	}

	// 尝试向下走
	if row+1 < rows && visited[row+1][col] != 1 && board[row+1][col] == word[index] {
		visited[row+1][col] = 1
		if backtrace(board, word, cols, rows, index+1, col, row+1, visited) {
			return true
		}
		visited[row+1][col] = 0
	}

	// 尝试向上走
	if row-1 >= 0 && visited[row-1][col] != 1 && board[row-1][col] == word[index] {
		visited[row-1][col] = 1
		if backtrace(board, word, cols, rows, index+1, col, row-1, visited) {
			return true
		}
		visited[row-1][col] = 0
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	fmt.Println(exist(board, "ABCCED"))
}
