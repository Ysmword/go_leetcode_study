package main

import "fmt"

// 方法：看八个位置，然后计算zero和one的值个数
func gameOfLife(board [][]int) {
	rl := len(board)
	postions := make([][]int, 0)
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			v := getValue(row, col, board, rl-1, len(board[row])-1)
			if v != board[row][col] {
				postions = append(postions, []int{row, col, v})
			}
		}
	}

	for _, p := range postions {
		board[p[0]][p[1]] = p[2]
	}
}

func getValue(row, col int, board [][]int, rl, cl int) int {
	postions := [][]int{
		{-1, -1}, // 左上角
		{-1, 0},  // 上
		{-1, 1},  // 右上角
		{0, -1},  // 左
		{0, 1},   // 右
		{1, -1},  // 右下角
		{1, 0},   // 下
		{1, 1},   // 右下角
	}
	countZero, countOne := 0, 0
	for _, p := range postions {
		rIndex, cIndex := row+p[0], col+p[1]
		if rIndex < 0 || rIndex > rl || cIndex < 0 || cIndex > cl {
			continue
		}
		if board[rIndex][cIndex] == 0 {
			countZero++
		}
		if board[rIndex][cIndex] == 1 {
			countOne++
		}
	}

	// 规则
	if board[row][col] == 1 {
		// 活细胞规则
		if countOne < 2 { // 如果活细胞周围八个位置的活细胞数少于两个，则该位置活细胞死亡；
			return 0
		}
		if countOne == 2 || countOne == 3 { // 如果活细胞周围八个位置有两个或三个活细胞，则该位置活细胞仍然存活；
			return 1
		}
		if countOne > 3 { // 如果活细胞周围八个位置有超过三个活细胞，则该位置活细胞死亡；
			return 0
		}
	}

	if board[row][col] == 0 {
		if countOne == 3 { // 如果死细胞周围正好有三个活细胞，则该位置死细胞复活；
			return 1
		}
	}
	return board[row][col]
}

// 优化方法：使用标记法，如果是0变1，为-1，如果是1变0则为-2，然后在根据这个规则替换值

func main() {
	board := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	fmt.Println(board)
	gameOfLife(board)
	fmt.Println(board)
}
