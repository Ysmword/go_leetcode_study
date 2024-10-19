package main

import (
	"fmt"

	"leetcode/common"
)

// 前序遍历变种
func hasPathSum(root *common.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var isHas func(*common.TreeNode, int) bool
	isHas = func(r *common.TreeNode, target int) bool {
		if r == nil {
			return false
		}
		target -= r.Val
		if target == 0 && r.Left == nil && r.Right == nil {
			return true
		}
		return isHas(r.Left, target) || isHas(r.Right, target)
	}
	return isHas(root, targetSum)
}

func main() {
	root := &common.TreeNode{Val: 5}
	node1 := &common.TreeNode{Val: 4}
	node2 := &common.TreeNode{Val: 8}
	node3 := &common.TreeNode{Val: 11}
	node4 := &common.TreeNode{Val: 13}
	node5 := &common.TreeNode{Val: 4}
	node6 := &common.TreeNode{Val: 7}
	node7 := &common.TreeNode{Val: 2}
	node8 := &common.TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node2.Left = node4
	node2.Right = node5
	node3.Left = node6
	node3.Right = node7
	node5.Right = node8
	fmt.Println(hasPathSum(root, 22))

	// node1 := &common.TreeNode{Val: 1}
	// node2 := &common.TreeNode{Val: 2}
	// node1.Left = node2
	// fmt.Println(hasPathSum(node1, 1))
}
