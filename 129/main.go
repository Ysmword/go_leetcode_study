package main

import (
	"fmt"
	"leetcode/common"
)

// 前序遍历的变种
func sumNumbers(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	dfs(root, 0, &sum)
	return sum
}

func dfs(root *common.TreeNode, cur int, sum *int) {
	if root == nil { // 左节点为空，右节点为空
		*sum += cur
		return
	}

	cur = cur*10 + root.Val
	if root.Left == nil && root.Right == nil { // 叶子节点
		*sum += cur
		return
	}

	if root.Left != nil {
		dfs(root.Left, cur, sum)
	}

	if root.Right != nil {
		dfs(root.Right, cur, sum)
	}
}

func main() {
	// root := &common.TreeNode{Val: 1}
	// node1 := &common.TreeNode{Val: 2}
	// node2 := &common.TreeNode{Val: 3}
	// root.Left = node1
	// root.Right = node2
	// fmt.Println(sumNumbers(root))

	root := &common.TreeNode{Val: 4}
	node1 := &common.TreeNode{Val: 9}
	node2 := &common.TreeNode{Val: 0}
	node3 := &common.TreeNode{Val: 5}
	node4 := &common.TreeNode{Val: 1}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	fmt.Println(sumNumbers(root))
}
