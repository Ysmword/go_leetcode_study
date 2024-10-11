package main

import (
	"fmt"
	"sort"
)

// 对Xstart进行升序排序，然后找交集，交集方式：Xend[i]>=Xstart[i+1], 要考虑下其中一个区间的最右边
func findMinArrowShots(points [][]int) int {
	sort.SliceStable(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	fmt.Println(points)
	first := points[0]
	count := 1
	for i := 1; i < len(points); i++ {
		if first[1] >= points[i][0] {
			first[1] = min(first[1], points[i][1])
			/*
				考虑这种情况：[0 9] [0 6] [2 9] [2 8] [3 9] [3 8] [3 9] [6 8] [7 12] [9 10]
				0 1 2 3 4 5 6 7 8 9 [0 9]
				0 1 2 3 4 5 6       [0 6]
					2 3 4 5 6 7 8 9 [2 9]
					2 3 4 5 6 7 8   [2 8]
					  3 4 5 6 7 8 9 [3 9]
					  3 4 5 6 7 8   [3 8]
					  3 4 5 6 7 8 9 [3 9]
					  3 4 5 6 7 8   [6 8]
					          7 8 9 10 11 12 [7 12]
							      9 10       [9 10]
				第一次取的是6或者3，第二次取9
			*/
		} else {
			first = points[i]
			count++
		}
	}
	return count
}

func main() {
	// fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	// fmt.Println(findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	// fmt.Println(findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	fmt.Println(findMinArrowShots([][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}))
}
