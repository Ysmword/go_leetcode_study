package main

import "fmt"

// 总体来说是：感染法，每重新感染一次就表示一个岛屿
// 感染法有两种方法：深度优先遍历 和 广度优先遍历

// 深度优先遍历
func numIslands(grid [][]byte) int {
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				dfs(grid, i, j)
			}
		}
	}
	return num
}

func dfs(grid [][]byte, row, col int) {
	if row < 0 || col < 0 || row >= len(grid) || col >= len(grid[row]) || grid[row][col] == '0' {
		return
	}
	grid[row][col] = '0'
	dfs(grid, row+1, col) // 右
	dfs(grid, row-1, col) // 左
	dfs(grid, row, col+1) // 下
	dfs(grid, row, col-1) // 上
}

// 广度优先遍历
type postion struct {
	row int
	col int
}

func numIslands1(grid [][]byte) int {
	q := make([]postion, 0)
	num := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				num++
				grid[i][j] = '0'
				q = append(q, postion{row: i, col: j})
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					if p.row-1 >= 0 && grid[p.row-1][p.col] == '1' {
						q = append(q, postion{row: p.row - 1, col: p.col})
						grid[p.row-1][p.col] = '0'
					}
					if p.row+1 < len(grid) && grid[p.row+1][p.col] == '1' {
						q = append(q, postion{row: p.row + 1, col: p.col})
						grid[p.row+1][p.col] = '0'
					}
					if p.col-1 >= 0 && grid[p.row][p.col-1] == '1' {
						q = append(q, postion{row: p.row, col: p.col - 1})
						grid[p.row][p.col-1] = '0'
					}
					if p.col+1 < len(grid[p.row]) && grid[p.row][p.col+1] == '1' {
						q = append(q, postion{row: p.row, col: p.col + 1})
						grid[p.row][p.col+1] = '0'
					}
				}
			}
		}
	}
	return num
}

func main() {
	num := numIslands1([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	})
	fmt.Println(num)
}
