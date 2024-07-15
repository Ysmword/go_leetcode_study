package main

import (
	"fmt"
	"reflect"
	"sort"
)

// 利用nums[k]=-(nums[x]+nums[y])，然后使用reflect.DeepEqual进行去重
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for k := 0; k < len(nums); k++ {
		target := -nums[k]
		x, y := 0, len(nums)-1
		for x < y {
			if x == k {
				x++
				continue
			}
			if y == k {
				y--
				continue
			}
			if nums[x]+nums[y] < target {
				x++
			} else if nums[x]+nums[y] > target {
				y--
			} else {
				vals := []int{nums[k], nums[x], nums[y]}
				sort.Ints(vals)
				exist := false
				for _, vs := range res {
					if reflect.DeepEqual(vals, vs) {
						exist = true
						break
					}
				}
				if !exist {
					res = append(res, vals)
				}
				x++
				y--
			}
		}
	}
	return res
}

// 进阶1：开始的时候过滤下，如果开始值重复出现，就不再继续后面计算
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	firstMap := make(map[int]struct{})
	for k := 0; k < len(nums); k++ {
		if _, ok := firstMap[nums[k]]; ok {
			continue
		}
		firstMap[nums[k]] = struct{}{}
		target := -nums[k]
		x, y := 0, len(nums)-1
		for x < y {
			if x == k {
				x++
				continue
			}
			if y == k {
				y--
				continue
			}
			if nums[x]+nums[y] < target {
				x++
			} else if nums[x]+nums[y] > target {
				y--
			} else {
				vals := []int{nums[k], nums[x], nums[y]}
				sort.Ints(vals)
				exist := false
				for _, vs := range res {
					if reflect.DeepEqual(vals, vs) {
						exist = true
						break
					}
				}
				if !exist {
					res = append(res, vals)
				}
				x++
				y--
			}
		}
	}
	return res
}

// 排序后，遍历过的，就不用在遍历+排序+二分查找
func threeSum2(nums []int) [][]int {
	sort.Ints(nums) // 排序是为了二分查找
	res := make([][]int, 0)
	for k := 0; k < len(nums)-1; k++ {
		if k > 0 && nums[k] == nums[k-1] { // 相同的值遍历过了，就不需要在遍历了
			continue
		}
		targes := -nums[k]
		x, y := k+1, len(nums)-1 // x从k+1开始，因为前面的都会被组合到
		for x < y {
			if nums[x]+nums[y] < targes { // 二分查找特性
				x++
				continue
			}
			if nums[x]+nums[y] > targes { // 二分查找特性
				y--
				continue
			}

			if x > k+1 && nums[x] == nums[x-1] { // 第二层不能在出现重复了
				x++
				continue
			}
			res = append(res, []int{nums[k], nums[x], nums[y]})
			x++
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum([]int{0, 1, 1}))
	fmt.Println(threeSum([]int{0, 0, 0}))

	fmt.Println(threeSum1([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum1([]int{0, 1, 1}))
	fmt.Println(threeSum1([]int{0, 0, 0}))

	fmt.Println(threeSum2([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum2([]int{0, 1, 1}))
	fmt.Println(threeSum2([]int{0, 0, 0}))
}
