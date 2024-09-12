package main

import (
	"fmt"
	"sort"
)

// 方法1: 直接排序取值
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 快排：取一边排序
func findKthLargest1(nums []int, k int) int {
	left := 0
	right := len(nums) - 1
	for {
		if left >= right { // 说明相遇了
			return nums[right]
		}
		index := partition(nums, left, right)
		if index+1 == k {
			return nums[index]
		} else if index+1 < k {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func partition(nums []int, left, right int) int {
	// 取最右边为基准
	base := nums[right] // 这里选的不好，可能会有最坏的情况
	for i := left; i < right; i++ {
		if nums[i] > base {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	return left
}

// 建堆
func findKthLargest2(nums []int, k int) int {

	heapSize := len(nums)
	buildMaxHeap(nums, heapSize)
	for i := len(nums) - 1; i >= len(nums)-k+1; i++ {
		nums[0], nums[i] = nums[i], nums[0]
		heapSize--
		maxHeapify(nums, 0, heapSize)
	}
	return nums[0]
}

func buildMaxHeap(nums []int, heapSize int) {
	for i := heapSize / 2; i >= 0; i-- {
		maxHeapify(nums, i, heapSize)
	}
}

func maxHeapify(nums []int, i, heapSize int) {
	l, r, largest := i*2+1, i*2+2, i
	if l < heapSize && nums[l] > nums[largest] {
		largest = l
	}
	if r < heapSize && nums[r] > nums[largest] {
		largest = r
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		maxHeapify(nums, largest, heapSize) // 注意这里是largest, 不是i, 因为largest已经被替换了, 所以要继续向下调整
	}
}

func main() {
	// fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	res := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(partition(res, 0, 5))
	fmt.Println(res)
}
