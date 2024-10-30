package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 根据搜索二叉树的特性，中序遍历得到的数组是有序的，然后计算相邻元素的差值，取最小值即可
// 最朴素的方法
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	vals := make([]int, 0)
	dfs(root, &vals)
	minDiff := math.MaxInt
	for i := 1; i < len(vals); i++ {
		minDiff = min(int(math.Abs(float64(vals[i-1])-float64(vals[i]))), minDiff)
	}
	return minDiff
}

func dfs(root *TreeNode, vals *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, vals)
	*vals = append(*vals, root.Val)
	dfs(root.Right, vals)
}

// 在中序遍历的过程中，直接计算相邻元素的差值，取最小值即可
func getMinimumDifference1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDiff := math.MaxInt
	prev := -1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != -1 {
			minDiff = min(minDiff, root.Val-prev) // 不需要转绝对值，因为中序遍历是升序的
		}
		prev = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return minDiff
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 0}
	node2 := &TreeNode{Val: 48}
	root.Left = node1
	root.Right = node2

	node3 := &TreeNode{Val: 12}
	node4 := &TreeNode{Val: 49}
	node2.Left = node3
	node2.Right = node4

	fmt.Println(getMinimumDifference(root))
}
