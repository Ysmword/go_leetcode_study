// 归并排序：主要思想：递归+合并；大概思路：分割左右数组，然后左边排序，右边排序，最后左右两边merge
package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2
	left := nums[0:mid]
	right := nums[mid:]
	newLefts := mergeSort(left)
	newRights := mergeSort(right)
	news := merge(newLefts, newRights)
	return news
}

func merge(left, right []int) []int {
	x, y := 0, 0
	news := make([]int, 0)
	for x < len(left) && y < len(right) {
		if left[x] < right[y] {
			news = append(news, left[x])
			x++
		} else {
			news = append(news, right[y])
			y++
		}
	}
	if x < len(left) {
		news = append(news, left[x:]...)
	}
	if y < len(right) {
		news = append(news, right[y:]...)
	}
	return news
}

func main() {
	nums := []int{9, 3, 10, 4, 7, 2, 1, 8, 8}
	fmt.Println(mergeSort(nums))
}
