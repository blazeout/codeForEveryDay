package _3_LRU缓存机制

type LRUCache struct {
	size     int
	capacity int
	//cache存储具体key的node, 以达到O(1)get值
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

// DLinkNode key, value
type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}

func Constructor(capacity int) LRUCache {
	head := &DLinkNode{}
	tail := &DLinkNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*DLinkNode),
		head:     head,
		tail:     tail,
	}
}

// Get 如果在map中就返回同时移动这个节点去头结点, 如果不存在就返回-1
func (c *LRUCache) Get(key int) int {
	if _, ok := c.cache[key]; !ok {
		return -1
	}
	c.moveToHead(c.cache[key])
	return c.cache[key].value
}

// 如果关键字 key 已经存在，则变更其数据值 value, 同时移动到最前面 moveToHead
// 如果不存在，则向缓存中插入该组 key-value addHead 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字 deleteTail。

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.cache[key]; !ok {
		node := &DLinkNode{
			key:   key,
			value: value,
		}
		c.cache[key] = node
		c.addHead(c.cache[key])
		c.size++
		if c.size > c.capacity {
			delete(c.cache, c.tail.prev.key)
			c.deleteTail(c.tail.prev)
			c.size--
		}
	}
	c.cache[key].value = value
	c.moveToHead(c.cache[key])
}

func (c *LRUCache) deleteNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) addHead(node *DLinkNode) {
	c.head.next.prev = node
	node.next = c.head.next
	node.prev = c.head
	c.head.next = node
}

func (c *LRUCache) moveToHead(node *DLinkNode) {
	// 包含 deleteNode 和 addHead
	c.deleteNode(node)
	c.addHead(node)
}

func (c *LRUCache) deleteTail(node *DLinkNode) {
	node.prev.next = c.tail
	c.tail.prev = node.prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */