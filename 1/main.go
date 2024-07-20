package main

import (
	"fmt"
)

// 排序+双指针，哈哈哈，行不通
// func twoSum(nums []int, target int) []int {
// 	sort.Ints(nums)
// 	fmt.Println(nums)
// 	x, y := 0, len(nums)-1
// 	for x < y {
// 		if nums[x]+nums[y] < target {
// 			x++
// 		} else if nums[x]+nums[y] > target {
// 			y--
// 		} else {
// 			return []int{x, y}
// 		}
// 	}
// 	return nil
// }

// 哈希表,一次遍历，反推法
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := numsMap[target-num]; ok {
			return []int{i, index}
		}
		numsMap[num] = i
	}
	return nil
}

// 暴力算法：双循环
func twoSum1(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
