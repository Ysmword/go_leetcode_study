package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 新链表，挑选元素 尾插法
func deleteDuplicates(head *ListNode) *ListNode {
	dummyNode := &ListNode{}
	dCur := dummyNode
	for n := head; n != nil; {
		cur := n.Next
		hasDup := false
		for cur != nil {
			if n.Val != cur.Val {
				break
			}
			hasDup = true
			cur = cur.Next
		}
		if hasDup {
			n = cur // 有相等的则不放进新的链表中
		} else {
			newNode := &ListNode{Val: n.Val}
			dCur.Next = newNode
			dCur = newNode
			n = n.Next
		}
	}
	return dummyNode.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// cur为跟后面不想等的值
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val // 用于记录val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

func main() {
	p1 := ListNode{Val: 1}
	p2 := ListNode{Val: 1}
	p1.Next = &p2
	p3 := ListNode{Val: 1}
	p2.Next = &p3
	p4 := ListNode{Val: 2}
	p3.Next = &p4
	p5 := ListNode{Val: 3}
	p4.Next = &p5

	printListNode(deleteDuplicates(&p1))

	l1 := ListNode{Val: 1}
	l2 := ListNode{Val: 2}
	l1.Next = &l2
	l3 := ListNode{Val: 3}
	l2.Next = &l3
	l4 := ListNode{Val: 3}
	l3.Next = &l4
	l5 := ListNode{Val: 4}
	l4.Next = &l5
	l6 := ListNode{Val: 4}
	l5.Next = &l6
	l7 := ListNode{Val: 5}
	l6.Next = &l7

	printListNode(deleteDuplicates(&l1))

	r1 := ListNode{Val: 1}
	r2 := ListNode{Val: 2}
	r1.Next = &r2
	r3 := ListNode{Val: 2}
	r2.Next = &r3
	printListNode(deleteDuplicates(&r1))
}

func printListNode(l *ListNode) {
	res := make([]int, 0)
	for n := l; n != nil; n = n.Next {
		res = append(res, n.Val)
	}
	fmt.Println(res)
}
