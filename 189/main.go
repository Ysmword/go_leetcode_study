package main

import (
	"fmt"
)

// 方法1：临时变量+取模+拷贝
func rotate(nums []int, k int) {
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		result[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, result)
}

// 方法2：冒泡排序翻版
func rotate1(nums []int, k int) {
	for n := 0; n < k; n++ {
		// 后往前遍历
		last := nums[len(nums)-1]
		for i := len(nums) - 1; i >= 1; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = last
	}
}

// 方法3：有效旋转次数为：a=k%n，也就代表a后的数字后提到前面，a的数字会提到后面;
// 这样需要有额外的空间
func rotate2(nums []int, k int) {
	result := make([]int, 0)
	result = append(result, nums[len(nums)-k%len(nums):]...)
	result = append(result, nums[0:len(nums)-k%len(nums)]...)
	copy(nums, result)
}

// 方法4：三次旋转，你一次全翻转，第二次对[0:k%len(nums)-1]翻转，第三次对[k%len(nums):]翻转
// 结合方法3看
func rotate3(nums []int, k int) {
	reverse(nums)
	reverse(nums[0 : k%len(nums)])
	reverse(nums[k%len(nums):])
}
func reverse(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for start, last := 0, len(nums)-1; start < last; {
		nums[start], nums[last] = nums[last], nums[start]
		start++
		last--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 3)
	fmt.Println(nums)
}
