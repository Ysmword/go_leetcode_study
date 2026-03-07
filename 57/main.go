package main

import "fmt"

/*
思路与算法
在给定的区间集合 X 互不重叠的前提下，当我们需要插入一个新的区间 S=[left,right] 时，我们只需要：
找出所有与区间 S 重叠的区间集合 X；
将 X′中的所有区间连带上区间 S 合并成一个大区间；
最终的答案即为不与 X′重叠的区间以及合并后的大区间。
*/

func insert(intervals [][]int, newInterval []int) [][]int {
	left, right := newInterval[0], newInterval[1]
	merged := false
	ans := [][]int{}
	for _, interval := range intervals {
		if interval[0] > right {
			// 在插入区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{left, right})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < left {
			// 在插入区间的左侧且无交集
			ans = append(ans, interval)
		} else {
			// 与插入区间有交集，计算它们的并集
			left = min(left, interval[0])
			right = max(right, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{left, right})
	}
	return ans
}

func main() {
	fmt.Println(insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))

	fmt.Println(insert1([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}

// 三阶段方法：第一步：先看 a[1] < new[0]，这个时候直接加入结果；第二步：a[0] < new[1],这个时候一定有交接,将合并的结果加入结果，第三步：添加剩余区间
func insert1(intervals [][]int, newInterval []int) [][]int {
	i, n := 0, len(intervals)-1
	res := make([][]int, 0)
	// 第一步：先看 a[1] < new[0]，这个时候直接加入结果
	for i <= n && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}

	// 第二步：a[0] < new[1],这个时候一定有交接
	for i <= n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(intervals[i][0], newInterval[0])
		newInterval[1] = min(intervals[i][1], newInterval[1])
		i++
	}
	res = append(res, newInterval)

	// 第三步：添加剩余区间
	for i <= n {
		res = append(res, intervals[i])
	}
	return res
}
