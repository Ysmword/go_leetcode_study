package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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

// Inorder 中序遍历
func (t *TreeNode) Inorder() {
	if t.Left != nil {
		t.Left.Inorder()
	}
	fmt.Println(t.Val)
	if t.Right != nil {
		t.Right.Inorder()
	}
}

func (t *TreeNode) Postorder() {
	if t.Left != nil {
		t.Left.Postorder()
	}
	if t.Right != nil {
		t.Right.Postorder()
	}
	fmt.Println(t.Val)
}
