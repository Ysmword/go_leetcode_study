package main

import (
	"container/heap"
	"fmt"
)

// 方法可行，但是时间复杂度太高，而且看着好像没看题呀
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	// 先组合
	var res [][]int
	for _, v := range nums1 {
		for _, vv := range nums2 {
			res = append(res, []int{v, vv})
		}
	}
	// 堆排序
	length := len(res)
	for i := length/2 - 1; i >= 0; i-- {
		adjustHeap(res, i, length)
	}
	temp := k
	for i := length - 1; i >= 0 && k > 0; i-- {
		res[i], res[0] = res[0], res[i]
		k--
		adjustHeap(res, 0, i)
	}
	return res[length-temp:]
}

func adjustHeap(res [][]int, i int, length int) {
	temp := res[i]
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if k+1 < length && res[k][0]+res[k][1] > res[k+1][0]+res[k+1][1] {
			k++
		}
		if res[k][0]+res[k][1] < temp[0]+temp[1] {
			res[i] = res[k]
			i = k
		}
	}
	res[i] = temp
}

func kSmallestPairs1(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }

func main() {
	nums1 := []int{1, 2, 4, 5, 6}
	nums2 := []int{3, 5, 7, 9}
	k := 3
	res := kSmallestPairs1(nums1, nums2, k)
	for _, v := range res {
		fmt.Println(v)
	}
}
