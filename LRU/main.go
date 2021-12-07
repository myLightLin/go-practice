package main

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func Constructor(capacity int) LRUCache {
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	return LRUCache{make(map[int]*LinkNode), capacity, head, tail}
}

func (this *LRUCache) Get(key int) int {
	cache := this.m
	if v, ok := cache[key]; ok {
		this.MoveToHead(v)
		return v.val
	} else {
		return -1
	}
}

func (this *LRUCache) RemoveNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) AddNode(node *LinkNode) {
	head := this.head
	// 从当前位置删除
	node.pre.next = node.next
	node.next.pre = node.pre
	// 移动到首位置
	node.next = head.next
	head.next.pre = node
	node.pre = head
	head.next = node
}

func (this *LRUCache) MoveToHead(node *LinkNode) {
	this.RemoveNode(node)
	this.AddNode(node)
}

func (this *LRUCache) Put(key int, value int) {
	head := this.head
	tail := this.tail
	cache := this.m
	if v, ok := cache[key]; ok {
		v.val = value
		this.MoveToHead(v)
	} else {
		v := &LinkNode{key, value, nil, nil}
		if len(cache) == this.cap {
			delete(cache, tail.pre.key)
			this.RemoveNode(tail.pre)
		}
		v.next = head.next
		v.pre = head
		head.next.pre = v
		head.next = v
		cache[key] = v
	}
}
