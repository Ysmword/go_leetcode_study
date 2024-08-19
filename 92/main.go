package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 一次遍历，先找到前驱点，并且使用头插法
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 示例：9-》7-》4-》5-》2-》3-》6，left=2，right=4
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode // 这个能处理left为1的情况
	for i := 0; i < left-1; i++ {
		pre = pre.Next // 找到7
	}

	cur := pre.Next // 找到4
	for i := 0; i < right-left; i++ {
		next := cur.Next     // 找到5
		cur.Next = next.Next // 找到4-》2
		next.Next = pre.Next // 5-》7
		pre.Next = next
	}
	return dummyNode.Next
}
