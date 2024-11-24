package main

import "fmt"

func searchRange1(nums []int, target int) []int {
	x, y := 0, len(nums)-1
	for x <= y {
		mid := (x + y) / 2
		if nums[mid] == target {
			left, right := mid, mid
			for left-1 >= 0 && nums[left-1] == target {
				left--
			}
			for right+1 < len(nums) && nums[right+1] == target {
				right++
			}
			return []int{left, right}
		}
		if nums[mid] > target {
			y = mid - 1
		} else {
			x = mid + 1
		}
	}
	return []int{-1, -1}
}

// 两次二分查找，分别查找左右边界
func searchRange(nums []int, target int) []int {
	x, y := 0, len(nums)-1
	// 先找左边界
	leftIdx := -1
	for x <= y {
		mid := (x + y) / 2
		if nums[mid] == target {
			leftIdx = mid
			y = mid - 1
		} else if nums[mid] > target {
			y = mid - 1
		} else {
			x = mid + 1
		}
	}

	rightIdx := -1
	x, y = 0, len(nums)-1
	for x <= y { // 找右边界
		mid := (x + y) / 2
		if nums[mid] == target {
			rightIdx = mid
			x = mid + 1
		} else if nums[mid] > target {
			y = mid - 1
		} else {
			x = mid + 1
		}
	}
	return []int{leftIdx, rightIdx}
}

func main() {
	fmt.Println(searchRange([]int{2, 2}, 2))
}
