package main

import (
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

func buildTree1(preorder []int, inorder []int) *common.TreeNode {
	// 1. 如果前序遍历为空，则直接返回
	if len(preorder) == 0 {
		return nil
	}

	// 2. 如果前序遍历只剩下一个元素，则直接返回节点
	root := &common.TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return root
	}

	// 3. 开始找元素在中序遍历的位置
	index := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}

	// 4. 切割中序遍历，以得到 左子树 和 右子树
	inLeft := inorder[0:index]
	inRight := inorder[index+1:]

	// 5. 切割前序遍历，以得到 左子树 和 右子树
	preLeft := preorder[1 : 1+len(inLeft)]
	preRight := preorder[1+len(inLeft):]

	// 6. 递归遍历 左子树 和 右子树
	root.Left = buildTree1(preLeft, inLeft)
	root.Right = buildTree1(preRight, inRight)
	return root
}

func main() {
	// t := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	// fmt.Println(t)

	t := buildTree1([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	t.Preorder()
	t.Inorder()
}
