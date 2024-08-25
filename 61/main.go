package main

import (
	"leetcode/common"
)

// 方法1：数组+数学计算
func rotateRight(head *common.ListNode, k int) *common.ListNode {
	lists := make([]*common.ListNode, 0)
	for node := head; node != nil; node = node.Next {
		lists = append(lists, node)
	}

	// 计算长度
	length := len(lists)
	newLists := make([]*common.ListNode, len(lists))
	for i := 0; i < len(lists); i++ {
		index := (i + k) % length
		lists[i].Next = nil
		newLists[index] = lists[i]
	}

	// 重新连接
	dummy := &common.ListNode{}
	cur := dummy
	for i := 0; i < len(newLists); i++ {
		cur.Next = newLists[i]
		cur = newLists[i]
	}
	return dummy.Next
}

// 方法2：闭合+数学计算，找到最后一个节点，然后断开即可
func rotateRight1(head *common.ListNode, k int) *common.ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}

	n := 1
	node := head
	for node.Next != nil {
		n++
		node = node.Next
	}

	// 旋转次数跟长度为倍数关系，则直接返回源指针
	if n%k == 0 {
		return head
	}

	lastIndex := n - k%n
	node.Next = head    // 形成闭环
	for lastIndex > 0 { // 找到最后一个点
		node = node.Next
		lastIndex--
	}
	ret := node.Next
	node.Next = nil
	return ret
}

func main() {
	head := common.CreateList([]int{1, 2, 3, 4, 5})
	res := rotateRight1(head, 2)
	common.PrintList(res)

	head = common.CreateList([]int{0, 1, 2})
	res = rotateRight1(head, 4)
	common.PrintList(res)
}
