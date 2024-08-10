package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 两数相加，利用取模和取余特性，再加个临时数
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	p := head
	carry := 0
	for l1 != nil && l2 != nil {
		val := l1.Val + l2.Val + carry
		carry = val / 10
		val = val % 10
		node := ListNode{
			Val: val,
		}
		p.Next = &node
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	p, carry = add(l1, p, carry)
	p, carry = add(l2, p, carry)
	if carry > 0 {
		p.Next = &ListNode{
			Val: carry,
		}
	}
	return head.Next
}

func add(l *ListNode, res *ListNode, carry int) (*ListNode, int) {
	for l != nil {
		val := l.Val + carry
		carry = val / 10
		val = val % 10
		node := ListNode{
			Val: val,
		}
		res.Next = &node
		res = res.Next
		l = l.Next
	}
	return res, carry
}

func main() {
	node1 := ListNode{
		Val: 2,
	}
	node2 := ListNode{
		Val: 4,
	}
	node3 := ListNode{
		Val: 3,
	}
	node1.Next = &node2
	node2.Next = &node3

	node4 := ListNode{
		Val: 5,
	}
	node5 := ListNode{
		Val: 6,
	}
	node6 := ListNode{
		Val: 4,
	}
	node4.Next = &node5
	node5.Next = &node6

	res := addTwoNumbers(&node1, &node4)
	printNode(res)
}

func printNode(l *ListNode) {
	res := make([]int, 0)
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	fmt.Println(res)
}
