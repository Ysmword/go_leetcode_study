package main

import (
	"fmt"
	"math"
)

// 方法1: 三指针
func findPeakElement1(nums []int) int {
	left, mid, right := -1, 0, 1
	for right <= len(nums) {
		leftVal := math.MinInt32
		midVal := nums[mid]
		rightVal := math.MinInt32
		if left != -1 {
			leftVal = nums[left]
		}
		if right != len(nums) {
			rightVal = nums[right]
		}

		if midVal > leftVal && midVal > rightVal {
			return mid
		}
		left++
		mid++
		right++
	}
	return 0
}

// 方法2: 寻找最大值
func findPeakElement(nums []int) int {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

// 方法3: 巧妙思想，人玩高处走，二分查找法
/*
如果 nums[i−1]<nums[i]>nums[i+1]，那么位置 i 就是峰值位置，我们可以直接返回 i 作为答案；

如果 nums[i−1]<nums[i]<nums[i+1]，那么位置 i 处于上坡，我们需要往右走，即 i←i+1；

如果 nums[i−1]>nums[i]>nums[i+1]，那么位置 i 处于下坡，我们需要往左走，即 i←i−1；

如果 nums[i−1]>nums[i]<nums[i+1]，那么位置 i 位于山谷，两侧都是上坡，我们可以朝任意方向走。

*/
func findPeakElement3(nums []int) int {
	get := func(i int) int {
		if i == -1 || i == len(nums) {
			return math.MinInt
		}
		return nums[i]
	}
	left, right := 0, len(nums)
	for {
		mid := (left + right) / 2
		if get(mid) > get(mid-1) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
