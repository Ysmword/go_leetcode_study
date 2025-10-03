package main

import "fmt"

// 数组是有序的，平方后最大值一定在数组的两端，所以我们可以使用双指针，从两端开始遍历
// 平方：负负得正，正正得正
func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	left, right := 0, len(nums)-1
	i := len(nums) - 1
	for left <= right {
		// 左边 大于 右边，取左边
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
			i--
			continue
		}
		// 左边 小于 右边，取右边
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res[i] = nums[right] * nums[right]
			right--
			i--
			continue
		}
		// 左边 等于 右边，取左边
		if nums[left]*nums[left] == nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
			i--
		}
	}
	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-4}))
}
