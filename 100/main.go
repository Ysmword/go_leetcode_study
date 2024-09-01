package main

import "leetcode/common"

// 同时遍历，同时比较；深度优先遍历
func isSameTree(p *common.TreeNode, q *common.TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	if p.Val != q.Val {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)
	return left && right
}

// 可以考虑广度优先遍历
