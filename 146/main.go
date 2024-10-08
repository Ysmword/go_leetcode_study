package main

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

// 双链表+map，map存储为key->node
type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key:   key,
		value: value,
	}
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: make(map[int]*DLinkedNode),
		head:  initDLinkedNode(0, 0),
		tail:  initDLinkedNode(0, 0),
		cap:   capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l

}

func (l *LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}

	node := l.cache[key]
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	if _, ok := l.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		l.cache[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.cap {
			removed := l.removeTail()
			delete(l.cache, removed.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}
