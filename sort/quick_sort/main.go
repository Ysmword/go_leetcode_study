// 快速排序：插入排序思想；但是又有点分治思想；因为你先找到一个基准点之后，这时候这个点就是排在她该在的位置了，这时候你在看左右两边
package main

import (
	"fmt"
)

// 看右边元素的适合排序的位置
func partition(arr []int, low, hight int) int {
	pivot := arr[hight] // 探讨这个元素所属的位置
	i := low            // 探讨元素能够所在的位置
	for j := 0; j < hight; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	} // 找到位置
	arr[i], arr[hight] = arr[hight], arr[i] // 将数据放到所属位置
	return i
}

func quickSortBasic(arr []int) {
	if len(arr) > 1 {
		pi := partition(arr, 0, len(arr)-1) // 先找位置
		quickSortBasic(arr[0:pi])           // 考虑左边
		quickSortBasic(arr[pi+1:])          // 考虑右边
	}
}

func quikSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSortBasic(arr)
}

func main() {
	nums := []int{1, 2, 3, 4, 3, 3, 3}
	quikSort(nums)
	fmt.Println(nums)
}
