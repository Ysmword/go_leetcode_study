package main

import (
	"fmt"
	"strings"
)

/*
             1
           1   1
         1   2   1
       1   3   3   1
     1   4   6   4   1
   1   5  10  10   5   1
*/

func generatePascalTriangle(n int) [][]int {
	if n <= 0 {
		return nil
	}
	triangle := make([][]int, n)
	triangle[0] = []int{1}

	for i := 1; i < n; i++ {
		prevRow := triangle[i-1]
		currentRow := make([]int, i+1)
		currentRow[0], currentRow[i] = 1, 1
		for j := 1; j < i; j++ {
			currentRow[j] = prevRow[j-1] + prevRow[j]
		}
		triangle[i] = currentRow
	}
	return triangle
}

func printPascalTriangle(triangle [][]int) {
	if len(triangle) == 0 {
		fmt.Println("空三角形")
		return
	}

	totalRows := len(triangle)
	for rowIdx, row := range triangle {
		// 计算当前行的缩进（每行缩进逐渐减少，实现居中）
		indent := strings.Repeat("  ", totalRows-rowIdx-1)
		fmt.Print(indent)

		// 打印当前行元素（固定宽度为 4，确保对齐）
		for _, num := range row {
			fmt.Printf("%4d", num)
		}
		fmt.Println() // 换行
	}
}

func main() {
	res := generatePascalTriangle(6)
	printPascalTriangle(res)
}
