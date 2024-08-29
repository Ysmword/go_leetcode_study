package main

import (
	"leetcode/common"
)

// 层级遍历
func maxDepth(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	depth := 0
	for len(queue) != 0 {
		depth++
		tempQueue := make([]*common.TreeNode, 0)
		for len(queue) != 0 { // 遍历一层
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
		}
		queue = tempQueue
	}
	return depth
}

// 广度优先遍历，不用创建数组
func maxDepth2(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*common.TreeNode, 0)
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}

// 递归遍历
func maxDepth1(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth1(root.Left) + 1
	right := maxDepth1(root.Right) + 1
	return max(left, right)
}
