package main

import (
	"fmt"

	"leetcode/common"
)

func buildTree(inorder []int, postorder []int) *common.TreeNode {
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}
	var dfsBuildTree func(iStart, iEnd, pStart, pEnd int) *common.TreeNode
	dfsBuildTree = func(iStart, iEnd, pStart, pEnd int) *common.TreeNode {
		if iStart > iEnd || pStart > pEnd {
			return nil
		}
		root := &common.TreeNode{Val: postorder[pEnd]}
		m := inorderMap[postorder[pEnd]]
		root.Left = dfsBuildTree(iStart, m-1, pStart, pStart+m-iStart-1)
		// pStart+m-iStart-1计算过程：后续数组的起始位置加上左子树长度-1 就是后后序数组结束位置了，左子树的长度 = 根节点索引-左子树
		root.Right = dfsBuildTree(m+1, iEnd, pStart+m-iStart, pEnd-1)
		// pStart+m-iStart计算过程：m-iStart表示左子树的长度
		return root
	}
	return dfsBuildTree(0, len(inorder)-1, 0, len(postorder)-1)
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	t := buildTree(inorder, postorder)
	t.Preorder()
	fmt.Println(t)
}
