package linkedlist

type CacheNode struct {
	key        int
	value      int
	next, prev *CacheNode
}

type LRUCache struct {
	capacity int
	left     *CacheNode
	right    *CacheNode
	cache    map[int]*CacheNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*CacheNode),
		left:     &CacheNode{},
		right:    &CacheNode{},
	}

	lru.left.next = lru.right
	lru.right.prev = lru.left
	return lru
}

func (this *LRUCache) remove(node *CacheNode) {
	prev, next := node.prev, node.next
	prev.next = next
	next.prev = prev
}

func (this *LRUCache) insert(node *CacheNode) {
	prev, next := this.right.prev, this.right
	prev.next = node
	next.prev = node
	node.next = next
	node.prev = prev
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.cache[key]
	if !ok {
		return -1
	}

	this.remove(v)
	this.insert(v)

	return v.value
}

func (this *LRUCache) Put(key, value int) {
	v, ok := this.cache[key]
	if ok {
		this.remove(v)
		v.value = value
		this.insert(v)
		return
	}

	if len(this.cache) == this.capacity {
		// capacity is full so we need to  remove least recently used from head of DL
		lru := this.left.next
		this.remove(lru)
		delete(this.cache, lru.key)
	}

	node := CacheNode{key: key, value: value}
	this.cache[key] = &node
	this.insert(&node)
	return
}
