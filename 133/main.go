package main

type Node struct {
	Val       int
	Neighbors []*Node
}

// 深度优先遍历，如果在遍历过程中已经创建过，就不用在创建了
func cloneGraph(node *Node) *Node {
	data := make(map[int]*Node)
	return dfs(node, data)
}

func dfs(node *Node, data map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	n, ok := data[node.Val]
	if ok {
		return n
	}
	n = &Node{Val: node.Val}
	data[node.Val] = n
	n.Neighbors = make([]*Node, len(node.Neighbors))
	for i := 0; i < len(node.Neighbors); i++ {
		n.Neighbors[i] = dfs(node.Neighbors[i], data)
	}
	return n
}
