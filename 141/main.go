package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 快慢指针
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
