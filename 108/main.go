package main

import (
	"fmt"

	"leetcode/common"
)

// 方法1: 递归, 平衡二叉搜索树，中序遍历就是升序排列，分治思想，中间计算
func sortedArrayToBST(nums []int) *common.TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *common.TreeNode {
	if left > right {
		return nil
	}
	val := (left + right) / 2
	root := &common.TreeNode{Val: nums[val]}
	root.Left = helper(nums, left, val-1)
	root.Right = helper(nums, val+1, right)
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	t := sortedArrayToBST(nums)
	if t == nil {
		fmt.Println("不可能为空，但是为空了")
		return
	}
	t.Inorder()
}
