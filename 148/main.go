package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 归并排序，链表跟数组的归并排序差不多，只是链表需要[快慢指针]遍历找到中间节点，然后断开
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next.Next
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil

	left := sortList(head)
	right := sortList(mid)

	return merge(left, right)
}

func merge(left *ListNode, right *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy

	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 3}
	l := sortList(head)

	printList(l)
}

func printList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
