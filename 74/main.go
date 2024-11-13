package main

import "fmt"

// 二次查找，第一次竖着查找，第二次横着查找
func searchMatrix(matrix [][]int, target int) bool {
	cx, cy := 0, len(matrix)-1
	for cx <= cy {
		mid := (cx + cy) / 2
		if matrix[mid][0] > target {
			cy = mid - 1
		} else if matrix[mid][0] < target {
			cx = mid + 1
		} else {
			return true
		}
	}
	row := cx - 1
	if row < 0 {
		return false
	}

	rx, ry := 0, len(matrix[0])-1
	for rx <= ry {
		mid := (rx + ry) / 2
		if matrix[row][mid] > target {
			ry = mid - 1
		} else if matrix[row][mid] < target {
			rx = mid + 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(matrix, 3))
}
