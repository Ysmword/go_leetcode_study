package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历一次求出长度，然后根据长度，和再次遍历，找到倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Val: -1}
	count := 0
	cur := dummyNode
	for n := head; n != nil; n = n.Next {
		count++
		cur.Next = n
		cur = cur.Next
	}

	if n == count { // 删除头的情况
		return dummyNode.Next.Next
	}

	index := count - n
	cur = dummyNode.Next
	fmt.Println(index)
	for i := 1; i <= count; i++ {
		if i == index {
			next := cur.Next
			cur.Next = next.Next
			i++
		}
		cur = cur.Next
	}
	return dummyNode.Next
}

// 方法2: 堆栈
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	nodes := make([]*ListNode, 0)
	dummyNode := &ListNode{Next: head}
	for n := dummyNode; n != nil; n = n.Next {
		nodes = append(nodes, n)
	}
	prev := nodes[len(nodes)-1-n]
	prev.Next = prev.Next.Next
	return dummyNode.Next
}
