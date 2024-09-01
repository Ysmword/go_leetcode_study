package main

import "leetcode/common"

// invertTree 翻转二叉树，前序遍历
func invertTree(root *common.TreeNode) *common.TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
