package main

import (
	"leetcode/common"
)

// 双指针
func partition(head *common.ListNode, x int) *common.ListNode {
	// 分为两个部分即可
	less := &common.ListNode{}
	greater := &common.ListNode{}
	lc, gc := less, greater
	for cur := head; cur != nil; {
		node := cur
		temp := cur.Next
		node.Next = nil

		if cur.Val < x {
			lc.Next = node
			lc = node
		} else {
			gc.Next = node
			gc = node
		}

		cur = temp
	}

	lc.Next = greater.Next
	return less.Next
}

func main() {
	list := common.CreateList([]int{1, 4, 3, 2, 5, 2})
	res := partition(list, 3)
	common.PrintList(res)

	list = common.CreateList([]int{2, 1})
	res = partition(list, 2)
	common.PrintList(res)

	list = nil
	res = partition(list, 0)
	common.PrintList(res)

	list1 := common.CreateList([]int{1, 4, 3, 0, 2, 5, 2})
	res1 := partition(list1, 3)
	common.PrintList(res1)

	list2 := common.CreateList([]int{1, 4, 3, 0, 5, 2})
	res2 := partition(list2, 2)
	common.PrintList(res2)

}
