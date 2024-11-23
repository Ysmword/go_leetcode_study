package main

import "fmt"

func search1(nums []int, target int) int {
	for index, v := range nums {
		if v == target {
			return index
		}
	}
	return -1
}

// 1. 二分查找，利用有序性做向右还是向左移动
func search(nums []int, target int) int {
	x, y := 0, len(nums)-1
	for x <= y {
		mid := (x + y) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 说明[0,mid]有序，这个等于0的情况是只有一或者两个个元素的情况
			if nums[0] <= target && target < nums[mid] {
				y = mid - 1
			} else {
				x = mid + 1
			}
		} else { // 说明[mid,len(nums)-1]有序
			if nums[mid] < target && target <= nums[len(nums)-1] {
				x = mid + 1
			} else {
				y = mid - 1
			}
		}
	}
	return -1
}

func main() {
	// fmt.Println(search1([]int{1, 2, 3, 4, 5}, 3))

	fmt.Println(search([]int{3, 1}, 1))

}
