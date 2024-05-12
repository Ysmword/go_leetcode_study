package main

import (
	"fmt"
)

/*
方法1removeElement1：双指针(交换法)，放后面
方法2removeElement2：双指针，放前面
方法3removeElement3：双指针，放前面，2和3的区别是减少了遍历的次数
*/

// 后移动
func backward(nums []int, index int) {
	if index == len(nums)-1 {
		nums[index] = -1
		return
	}

	for i := index; i < len(nums)-1; i++ {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}
}

func removeElement(nums []int, val int) int {
	x, tail, count := 0, len(nums)-1, 0
	for x <= tail {
		if nums[x] == val {
			if nums[tail] == val {
				count++
				tail--
				continue
			}
			if nums[tail] != val {
				nums[x], nums[tail] = nums[tail], nums[x]
				x++
				tail--
				count++
			}
			continue
		}
		x++
	}
	return len(nums) - count
}

func removeElement1(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if v != val {
			nums[left] = v
			left++
		}
	}
	return left
}

func removeElement2(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left // 为什么直接取left，因为left部分放置就是不等于val的值
}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement1(nums, 2))
}
