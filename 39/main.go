package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	backtrack(0, candidates, target, 0, []int{}, &res)
	return res
}

func backtrack(start int, candidates []int, target, state int, track []int, res *[][]int) {
	// 终止条件
	if state == target {
		// 拷贝一份
		track2 := make([]int, len(track))
		copy(track2, track)
		*res = append(*res, track2)
		return
	}

	for i := start; i < len(candidates); i++ {
		if state+candidates[i] > target { // 因为数组已经排序过，后面的数都大于target，就不用再继续了
			break
		}
		track = append(track, candidates[i])
		backtrack(i, candidates, target, state+candidates[i], track, res)
		// 这里为什么使用i，是因为可以重复选择同一个数
		// 这里为什么不用start，因为start的值不会变，会导致有重复的值
		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
