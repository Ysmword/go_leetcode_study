package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	arr []int
}

// 将中序遍历的结果求出来，然后根据next和hasNext来实现迭代器
func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{
		arr: make([]int, 0),
	}
	inorder(root, &b.arr)
	return b
}

func inorder(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, arr)
	*arr = append(*arr, root.Val)
	inorder(root.Right, arr)
}

func (b *BSTIterator) Next() int {
	val := b.arr[0]
	b.arr = b.arr[1:]
	return val
}

func (b *BSTIterator) HasNext() bool {
	return len(b.arr) != 0
}
