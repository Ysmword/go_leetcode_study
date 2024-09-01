package main

import "leetcode/common"

// 对称二叉树， 递归
/*
	p指针和q指针一开始都指向这棵树的根，随后 p 右移时，q 左移，p 左移时，
	q 右移。每次检查当前 p 和 q 节点的值是否相等，如果相等再判断左右子树是否对称。
*/
func isSymmetric(root *common.TreeNode) bool {
	return check(root, root)
}

func check(p, q *common.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 迭代方法
func isSymmetric1(root *common.TreeNode) bool {
	u, v := root, root
	q := make([]*common.TreeNode, 0)
	q = append(q, u)
	q = append(q, v)

	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
