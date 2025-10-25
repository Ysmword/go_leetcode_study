package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 平衡二叉树：左右子树高度差不超过1
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := isBalanced(root.Left)
	right := isBalanced(root.Right)

	// 左右子树都平衡，且左右子树高度差不超过1
	if left && right && math.Abs(float64(getNodeHeight(root.Left)-getNodeHeight(root.Right))) <= 1 {
		return true
	}
	return false
}

func getNodeHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getNodeHeight(root.Left)
	right := getNodeHeight(root.Right)
	return max(left, right) + 1
}
