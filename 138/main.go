package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 哈希表+递归
func copyRandomList(head *Node) *Node {
	cacheNode := make(map[*Node]*Node)
	return deepCopy(head, cacheNode)
}

func deepCopy(head *Node, cacheNode map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}
	if n, ok := cacheNode[head]; ok { // 已经遍历过或者说是创建过，就不用在创建了
		return n
	}
	newNode := &Node{Val: head.Val}
	cacheNode[head] = newNode
	newNode.Next = deepCopy(head.Next, cacheNode)
	newNode.Random = deepCopy(head.Random, cacheNode)
	return newNode
}

// 迭代+节点拆分
func copyRandomList1(head *Node) *Node {
	if head == nil {
		return nil
	}

	// 遍历node
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}

	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next // 连接原来
		if nodeNew.Next != nil {
			nodeNew.Next = nodeNew.Next.Next // 连接新
		}
	}
	return headNew
}
