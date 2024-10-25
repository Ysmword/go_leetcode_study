package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val { // 如果等于某一个值，表示我已经找到值了
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)   // 在左子树找, 找到的话就返回
	right := lowestCommonAncestor(root.Right, p, q) // 在右子树找, 找到的话就返回
	if left != nil && right != nil {                // 如果左右子树都找到了, 说明当前节点就是最近公共祖先
		return root
	}
	if left == nil { // 如果左子树没找到, 就返回右子树，因为题意：p 和 q 均存在于给定的二叉树中。
		return right
	}
	return left // 如果右子树没找到, 就返回左子树，因为题意：p 和 q 均存在于给定的二叉树中。
}
