package main

import "fmt"

// 就像整理扑克牌：从第二个元素开始，每次将一个元素插入到前面已排序序列的正确位置。
func insertSort(arr []int) {

	for i := 1; i < len(arr); i++ {
		key := arr[i] // 要插入的元素
		j := i - 1    // 从i-1开始向前比较
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // 往后移动
			j--
		}
		arr[j+1] = key
	}
}

func main() {
	nums := []int{5, 2, 4, 6, 1, 3}
	insertSort(nums)
	fmt.Println(nums)
}
