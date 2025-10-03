package main

import "fmt"

// 实现一个优先队列
type PriorityQueue struct {
	queue []int
}

func (p *PriorityQueue) Push(val int) {
	for !p.Empty() && val > p.queue[len(p.queue)-1] { // 不能取等号，否则会缺失数据，例子：-7, -8, 7, 5, 7, 1, 6, 0；k=4
		p.queue = p.queue[:len(p.queue)-1] // 全出去，因为有面的元素更大
	}
	// 为什么小的元素要留着？因为如果队头的元素已经滑出窗口，得保证后续的元素能成为队头
	p.queue = append(p.queue, val) // 加入新元素
}

func (p *PriorityQueue) Pop(item int) {
	// 弹出队头元素
	val := p.queue[0]
	if val == item {
		p.queue = p.queue[1:]
	}
}

func (p *PriorityQueue) Front() int {
	return p.queue[0]
}

func (p *PriorityQueue) Empty() bool {
	return len(p.queue) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	pq := &PriorityQueue{}
	for i := 0; i < k && i < len(nums); i++ {
		pq.Push(nums[i])
	}

	res := make([]int, 0)
	res = append(res, pq.Front())
	for i := k; i < len(nums); i++ {
		pq.Pop(nums[i-k])
		pq.Push(nums[i])
		res = append(res, pq.Front())
	}

	fmt.Println(pq.queue)
	return res
}

func main() {
	fmt.Println(maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4))
}
