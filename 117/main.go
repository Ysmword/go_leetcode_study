package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 层级遍历的变种，使用队列
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}
	for len(queue) != 0 {
		levelSize := len(queue)
		nextQueue := make([]*Node, 0)
		current := queue[0]
		queue = queue[1:]
		for i := 0; i < levelSize; i++ {
			if current.Left != nil {
				nextQueue = append(nextQueue, current.Left)
			}
			if current.Right != nil {
				nextQueue = append(nextQueue, current.Right)
			}
			if len(queue) != 0 {
				node := queue[0]
				queue = queue[1:]
				current.Next = node
				current = node
			}
		}
		queue = nextQueue
	}
	return root
}

func main() {
	root := &Node{
		Val: 1,
	}
	n1 := &Node{
		Val: 2,
	}
	n2 := &Node{
		Val: 3,
	}
	root.Left = n1
	root.Right = n2

	n3 := &Node{
		Val: 4,
	}
	n4 := &Node{
		Val: 5,
	}
	n1.Left = n3
	n1.Right = n4

	n5 := &Node{
		Val: 5,
	}
	n2.Right = n5

	connect(root)
}
