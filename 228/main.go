package main

import (
	"fmt"
	"strconv"
)

// 区间:双指针
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	res := make([]string, 0)
	start, end := 0, 1
	for end < len(nums) { // 无法看到最后一位
		if nums[end]-nums[end-1] != 1 {
			if end-start != 1 {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
			} else {
				res = append(res, strconv.Itoa(nums[start]))
			}
			start = end
			end = end + 1
		} else {
			end++
		}
	}
	if nums[end-1]-nums[end-2] != 1 { // 包含最后一位
		res = append(res, strconv.Itoa(nums[end-1]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[end-1]))
	}
	return res
}

func main() {
	fmt.Println(summaryRanges([]int{0, 1, 2, 4, 5}))
	fmt.Println(summaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
}
