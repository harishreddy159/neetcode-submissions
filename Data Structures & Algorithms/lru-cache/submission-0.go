type LRUCache struct {
   	Capacity int
	Cache    map[int]*Node
	Head *Node
	Tail *Node 
}

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
    return LRUCache{
		Capacity: capacity,
		Cache: make(map[int]*Node),
		Head: head,
		Tail: tail,
	}
}

func (lru *LRUCache) Get(key int) int {
   	node, ok := lru.Cache[key]
	if !ok {
		return -1
	}

	lru.Remove(node)
	lru.Insert(node)

	return node.Value 
}

func (lru *LRUCache) Put(key int, value int) {
    if node, ok := lru.Cache[key]; ok {
		node.Value = value
		lru.Remove(node)
		lru.Insert(node)
		return
	}

	newNode := &Node{
		Key:   key,
		Value: value,
	}
	if len(lru.Cache) == lru.Capacity {
		lruNode := lru.Head.Next
		lru.Remove(lruNode)
		delete(lru.Cache, lruNode.Key)
	}
	lru.Insert(newNode)
	lru.Cache[key] = newNode
}

// Removes a node from the doubly linked list.
func (lru *LRUCache) Remove(node *Node) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev

	node.Next, node.Prev = nil, nil
}

// Inserts a node just before the tail
// (marks it as most recently used).
func (lru *LRUCache) Insert(node *Node) {
	prev := lru.Tail.Prev
	prev.Next = node
	node.Prev = prev
	node.Next = lru.Tail
	lru.Tail.Prev = node
}
