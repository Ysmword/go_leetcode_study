package main

import (
	"fmt"
	"sort"
)

/*
merge：非递减特性和数组元素后移思想
merge1：拷贝后，排序
merge2：非递减特性和额外空间
merge3：非递减特性和逆向双指针特性
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	x, y := 0, 0 // x,y分别表示nums1和nums2的下表
	for y < n {  // 只要将n个元素合入nums1即可
		if nums1[x] <= nums2[y] && m != 0 { // 下面移动nums1中的元素
			// m==0表示移动了m个元素，
			x++
			m--
			continue
		}
		// 下面移动nums2中的元素
		move(nums1, x, nums2[y])
		x++
		y++
	}
}

func move(nums []int, targetIndex int, value int) {
	for i := len(nums) - 2; i >= targetIndex; i-- {
		nums[i+1] = nums[i]
	}
	nums[targetIndex] = value
}

// merge1 拷贝后，排序
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2) // nums1[m:] 表示下标m之后的所有元素
	sort.Ints(nums1)
}

// merge2 非递减特性和额外空间
func merge2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	x, y := 0, 0
	for x < m && y < n {
		if nums1[x] <= nums2[y] {
			sorted = append(sorted, nums1[x])
			x++
			continue
		}
		sorted = append(sorted, nums2[y])
		y++
	}
	if x < m {
		sorted = append(sorted, nums1[x:]...)
	}
	if y < n {
		sorted = append(sorted, nums2[y:]...)
	}
	copy(nums1, sorted)
}

// merge3 非递减特性和额外空间
func merge3(nums1 []int, m int, nums2 []int, n int) {
	tail := len(nums1) - 1
	x, y := m-1, n-1
	for x >= 0 && y >= 0 {
		if nums1[x] >= nums2[y] {
			nums1[tail] = nums1[x]
			x--
			tail--
			continue
		}

		nums1[tail] = nums2[y]
		tail--
		y--
	}

	for x >= 0 {
		nums1[tail] = nums1[x]
		x--
		tail--
	}

	for y >= 0 {
		nums1[tail] = nums2[y]
		tail--
		y--
	}
}

func main() {
	nums1 := []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	nums2 := []int{1, 2, 2}
	merge(nums1, 6, nums2, 3)
	fmt.Println("0", nums1)

	nums1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	nums2 = []int{1, 2, 2}
	merge1(nums1, 6, nums2, 3)
	fmt.Println("1", nums1)

	nums1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	nums2 = []int{1, 2, 2}
	merge2(nums1, 6, nums2, 3)
	fmt.Println("2", nums1)

	nums1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	nums2 = []int{1, 2, 2}
	merge3(nums1, 6, nums2, 3)
	fmt.Println("3", nums1)

	var l int64 = 4
	var c int64 = 7
	a := make([]int, l, c)
	fmt.Println(a)

	fmt.Println((256 * 3) >> 2)
}
