package main

import (
	"fmt"
	"sort"
)

// 方法1：排序后，取中间位置
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 方法2：数量统计后比取大于len(nums)/2的数字
func majorityElement1(nums []int) int {
	cache := make(map[int]int)
	for _, num := range nums {
		cache[num]++
		if cache[num] > len(nums)/2 {
			return num
		}
	}
	return 0 // 不可能找不到
}

// 方法3：摩尔投票，正负抵消，最后假设的数字一定是数量最多的数字
// 由于题目要求，一定会存在，因此最后假设的数字一定符合，如果题目没有要求，则需要统计验证下
func majorityElement3(nums []int) int {
	result, vote := 0, 0 // result假设的结果，vote投票数量
	for _, num := range nums {
		if vote == 0 { // 如果初始投票数量为0，就开始假设当前数是否符合
			result = num
			vote++
			continue
		}
		if result != num { // 不是就减票
			vote--
			continue
		}
		vote++ // 是则加票
	}
	return result
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majorityElement3(nums))
}
