package main

import "fmt"

// 对排序，实现流程：第一步建立大堆，第二步排序
// 大堆满足：根节点大于左右子节点

// n 表示要调整的结束位置，i表示当前节点
func buildHeap(arr []int, n, i int) {
	largest := i
	leftIndex := 2*i + 1
	rightIndex := 2*i + 2

	if leftIndex < n && arr[leftIndex] > arr[largest] {
		largest = leftIndex
	}

	if rightIndex < n && arr[rightIndex] > arr[largest] {
		largest = rightIndex
	}

	if largest != i {
		arr[largest], arr[i] = arr[i], arr[largest]
		buildHeap(arr, n, largest) // 交换之后，要看下largest的节点是否满足大根堆，如果不满足，则继续调整
	}
}

func heapSort(arr []int) {
	n := len(arr)

	// 1.建立大根堆
	for i := n/2 - 1; i >= 0; i-- { // 从最后一个节点给他建立大根堆
		buildHeap(arr, n, i)
	}

	// 2. 排序
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		buildHeap(arr, i, 0) // 从新建堆，为什么从0开始，因为这个时候只有根节点不满足，前提都已经满足了
	}
}

func main() {
	arr := []int{10, 5, 8, 4, 1, 6}
	heapSort(arr)
	fmt.Println(arr)
}
