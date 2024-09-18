package main

type MinStack struct {
	stack    []int // 记录栈
	minStack []int // 记录出现过的最小值, 由于后进先出特性，能保证每次pop都可以判断下是否需要从minStack中pop
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	if len(m.minStack) == 0 {
		m.minStack = append(m.minStack, val)
	} else {
		if val <= m.minStack[len(m.minStack)-1] {
			m.minStack = append(m.minStack, val)
		}
	}
}

func (m *MinStack) Pop() {
	val := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if val == m.minStack[len(m.minStack)-1] {
		m.minStack = m.minStack[:len(m.minStack)-1]
	}
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	return m.minStack[len(m.minStack)-1]
}
