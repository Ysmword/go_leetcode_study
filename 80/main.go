package main

import "fmt"

// 方法1：临时变量+限制出现次数
// 方法2：快慢双指针

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	tempVal, tempValCount, count, left := nums[0]-1, 0, 2, 0 // 临时变量(初始值一定要保证不能等于最小值)，临时变量出现次数，出现次数，从左到右指针
	for _, num := range nums {
		if tempVal != num { // 初次遇见
			tempValCount = 1
			nums[left] = num
			tempVal = num
			left++
			continue
		}
		if tempVal == num && tempValCount == 1 { // 说明第二次遇见
			tempValCount++
			nums[left] = num
			left++
			continue
		}
		if tempVal == num && tempValCount >= count { // 第三次或者以上遇见
			tempValCount++
			continue
		}
	}
	return left
}

func removeDuplicates1(nums []int) int {
	if len(nums) <= 2 { // 如果长度小于等于2，一定不会出现重复出现2次以上的数字
		return len(nums)
	}
	slow, fast := 2, 2 // slow表示保留位置，fast表示遍历数组的下标
	for fast < len(nums) {
		if nums[fast] != nums[slow-2] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(removeDuplicates1(nums))
	fmt.Println(nums)
}
