package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 分层遍历，求每层的平均值
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	result := make([]float64, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	for len(queue) > 0 {
		lenght := len(queue)
		sum := 0.0
		for i := 0; i < lenght; i++ {
			node := queue[0]
			sum += float64(node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, sum/float64(lenght))
	}
	return result
}
