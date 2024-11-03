package main

import "fmt"

type postion struct {
	row int
	col int
}

// 正常的广度优先遍历
func solve(board [][]byte) {

	q := make([]postion, 0)
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if visited[i][j] {
				continue
			}
			needSetX := true
			needSetXPos := make([]postion, 0)
			if board[i][j] == 'O' {
				q = append(q, postion{row: i, col: j})
				needSetXPos = append(needSetXPos, postion{row: i, col: j})
				visited[i][j] = true
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					if p.row == 0 || p.row == len(board)-1 || p.col == 0 || p.col == len(board[i])-1 {
						if board[p.row][p.col] == 'O' {
							needSetX = false // 碰到边界说明不可能被包围
						}
					}
					if p.row-1 >= 0 && board[p.row-1][p.col] == 'O' && !visited[p.row-1][p.col] {
						q = append(q, postion{row: p.row - 1, col: p.col})
						visited[p.row-1][p.col] = true
						needSetXPos = append(needSetXPos, postion{row: p.row - 1, col: p.col})
					}
					if p.row+1 < len(board) && board[p.row+1][p.col] == 'O' && !visited[p.row+1][p.col] {
						q = append(q, postion{row: p.row + 1, col: p.col})
						visited[p.row+1][p.col] = true
						needSetXPos = append(needSetXPos, postion{row: p.row + 1, col: p.col})
					}
					if p.col-1 >= 0 && board[p.row][p.col-1] == 'O' && !visited[p.row][p.col-1] {
						q = append(q, postion{row: p.row, col: p.col - 1})
						visited[p.row][p.col-1] = true
						needSetXPos = append(needSetXPos, postion{row: p.row, col: p.col - 1})
					}
					if p.col+1 < len(board[i]) && board[p.row][p.col+1] == 'O' && !visited[p.row][p.col+1] {
						q = append(q, postion{row: p.row, col: p.col + 1})
						visited[p.row][p.col+1] = true
						needSetXPos = append(needSetXPos, postion{row: p.row, col: p.col + 1})
					}
				}
			}
			if needSetX && len(needSetXPos) > 0 {
				setX(board, needSetXPos)
			}
		}
	}
}

func setX(board [][]byte, pos []postion) {
	for _, p := range pos {
		board[p.row][p.col] = 'X'
	}
}

// 后面优化，反向思维：搜集边界上的O，然后对边界上的O进行遍历

func main() {
	board := [][]byte{
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	for _, b := range board {
		fmt.Println(string(b))
	}
	solve(board)
	fmt.Println("========")
	for _, b := range board {
		fmt.Println(string(b))
	}
}
