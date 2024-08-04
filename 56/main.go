package main

import (
	"fmt"
	"sort"
)

// 按左端点排序，排序后遍历合并，注意最后一位，算法相对复杂
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	i := 0
	for ; i < len(intervals)-1; i++ {
		if intervals[i][1] < intervals[i+1][0] {
			res = append(res, intervals[i])
		} else {
			if intervals[i][1] > intervals[i+1][1] {
				intervals[i+1] = intervals[i]
			} else {
				intervals[i+1][0] = intervals[i][0]
			}
		}
	}

	// 看最后一位
	if intervals[i][1] > intervals[i-1][0] {
		res = append(res, intervals[i])
	} else {
		if intervals[i][1] < intervals[i-1][1] {
			intervals[i] = intervals[i-1]
		}
		res = append(res, intervals[i])
	}
	return res
}

// 左端点排序，利用临时变量处理边界问题
func merge1(intervals [][]int) [][]int {
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	for _, interval := range intervals {
		m := len(res)
		if m > 0 && interval[0] <= res[m-1][1] { // 可以合并
			res[m-1][1] = max(interval[1], res[m-1][1])
		} else { // 不能合并
			res = append(res, interval)
		}
	}
	return res
}

func main() {
	fmt.Println(merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))

	fmt.Println(merge([][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
		{1, 10},
	}))

	fmt.Println(merge1([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}))

	fmt.Println(merge1([][]int{
		{2, 3},
		{4, 5},
		{6, 7},
		{8, 9},
		{1, 10},
	}))
}
