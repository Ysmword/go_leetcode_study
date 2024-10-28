package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 翻转｜逆向插入
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	queue := []*TreeNode{root}
	result := make([][]int, 0)
	level := 1
	for len(queue) > 0 {
		res := make([]int, 0)
		length := len(queue)
		tq := make([]*TreeNode, 0)
		if level%2 == 1 {
			for i := 0; i < length; i++ {
				node := queue[0]
				queue = queue[1:]
				res = append(res, node.Val)
				if node.Left != nil {
					tq = append(tq, node.Left)
				}
				if node.Right != nil {
					tq = append(tq, node.Right)
				}
			}
		} else {
			for i := length - 1; i >= 0; i-- {
				node := queue[i]
				queue = queue[:i]
				res = append(res, node.Val)
				if node.Right != nil {
					tq = append([]*TreeNode{node.Right}, tq...) // 这里要反着加入
				}
				if node.Left != nil {
					tq = append([]*TreeNode{node.Left}, tq...)
				}
			}
		}
		queue = tq
		result = append(result, res)
		level++
	}
	return result
}
