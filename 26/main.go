package main

import "fmt"

/*
	方法1：双指针，唯一值放前面
*/

func removeDuplicates(nums []int) int {
	temp, left := 0, 0
	for i, v := range nums {
		if i == 0 {
			temp = v
			left++
			continue
		}
		if v != temp {
			nums[left] = v
			temp = v
			left++
		}
	}
	return left
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
