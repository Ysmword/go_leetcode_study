package main

import (
	"fmt"

	"leetcode/common"
)

// 分治思想，要知道前序遍历和中序遍历的特点
func buildTree(preorder []int, inorder []int) *common.TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	return dfsBuildTree(preorder, inorderMap, 0, 0, len(inorder)-1)
}

func dfsBuildTree(preorder []int, inorderMap map[int]int, i, l, r int) *common.TreeNode {
	if l > r {
		return nil
	}
	root := &common.TreeNode{Val: preorder[i]}
	m := inorderMap[preorder[i]]
	root.Left = dfsBuildTree(preorder, inorderMap, i+1, l, m-1)
	root.Right = dfsBuildTree(preorder, inorderMap, i+m-l+1, m+1, r)
	return root
}

func main() {
	t := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	fmt.Println(t)
}
