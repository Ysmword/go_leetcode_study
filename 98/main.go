package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
func isValidBST(root *TreeNode) bool {
	data := make([]int, 0)
	return dfs(root, &data)
}

func dfs(root *TreeNode, data *[]int) bool {
	if root == nil {
		return true
	}
	left := dfs(root.Left, data)
	if len(*data) > 0 && root.Val <= (*data)[len(*data)-1] {
		return false
	}
	*data = append(*data, root.Val)
	right := dfs(root.Right, data)
	return left && right
}
