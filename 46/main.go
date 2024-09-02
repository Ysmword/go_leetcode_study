package main

import "fmt"

// 回溯+枝剪
func permute(nums []int) [][]int {
	choice := make(map[int]struct{})
	res := make([][]int, 0)
	backtrack(nums, []int{}, choice, &res)
	return res
}

func backtrack(nums []int, track []int, choice map[int]struct{}, res *[][]int) {
	if len(track) == len(nums) {
		temp := make([]int, len(track))
		copy(temp, track)
		*res = append(*res, temp)
		return
	}
	for i := 0; i < len(nums); i++ {
		if _, ok := choice[nums[i]]; ok {
			continue
		}
		choice[nums[i]] = struct{}{}
		track = append(track, nums[i])
		backtrack(nums, track, choice, res)
		track = track[:len(track)-1]
		delete(choice, nums[i])
	}
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{0, 1}))
}
