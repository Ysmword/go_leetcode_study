package main

import "fmt"

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 使用栈
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var prev *TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}
		left, right := curr.Left, curr.Right
		if right != nil {
			stack = append(stack, right) // 先将右节点压入栈中
		}
		if left != nil {
			stack = append(stack, left) // 左节点不为空，则压入栈中
		}
		prev = curr // 当前节点赋值给prev，用于下一轮循环的连接操作
	}
}

// Preorder 前序遍历
func (t *TreeNode) Preorder() {
	fmt.Println(t.Val)
	if t.Left != nil {
		t.Left.Preorder()
	}
	if t.Right != nil {
		t.Right.Preorder()
	}
}

func main() {
	root := &TreeNode{Val: 1}
	node1 := &TreeNode{Val: 2}
	node2 := &TreeNode{Val: 5}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 6}
	root.Left = node1
	root.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Right = node5
	flatten(root)
	root.Preorder()
}
