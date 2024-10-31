package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历，找到第k个元素，返回
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	val := 0
	var dfs func(root *TreeNode, k *int)
	dfs = func(root *TreeNode, k *int) {
		if root == nil {
			return
		}
		dfs(root.Left, k)
		*k--
		if *k == 0 {
			val = root.Val
			return
		}
		dfs(root.Right, k)
	}
	dfs(root, &k)
	return val
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 2}
	fmt.Println(kthSmallest(root, 1))
}

// 抽空想下，如何通过迭代法实现
