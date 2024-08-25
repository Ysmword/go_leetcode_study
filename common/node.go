package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// create list node
func CreateList(data []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for i := 0; i < len(data); i++ {
		node := &ListNode{Val: data[i]}
		cur.Next = node
		cur = node
	}
	return dummy.Next
}

func PrintList(head *ListNode) {
	data := make([]int, 0)
	for node := head; node != nil; node = node.Next {
		data = append(data, node.Val)
	}
	fmt.Println(data)
}
