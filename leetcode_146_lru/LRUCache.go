package leetcode_146_lru

type LRUCache struct {
	size       int
	cap        int
	cache      map[int]*dLinkCache
	head, tail *dLinkCache
}

type dLinkCache struct {
	key, value int
	prev, next *dLinkCache
}

func initDLinkCache(key, value int) *dLinkCache {
	return &dLinkCache{
		key:   key,
		value: value,
	}
}

// Constructor 最近最少使用
func Constructor(cap int) LRUCache {
	cache := LRUCache{
		cap:   cap,
		cache: make(map[int]*dLinkCache, 0),
		head:  initDLinkCache(0, 0),
		tail:  initDLinkCache(0, 0),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (l LRUCache) Get(key int) int {
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	//获取到当前的key
	node := l.cache[key]
	//移动node到链表头部
	l.moveToHead(node)
	return node.value
}

func (l LRUCache) Put(key, val int) {
	if _, ok := l.cache[key]; !ok {
		node := initDLinkCache(key, val)
		l.cache[key] = node
		l.addNode(node)
		l.size++
		if l.size > l.cap {
			remove := l.removeTail()
			delete(l.cache, remove.key)
			l.size--
		}
	} else {
		node := l.cache[key]
		node.value = val
		l.moveToHead(node)
	}
}

func (l LRUCache) removeNode(node *dLinkCache) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// addNode 增加新节点
func (l LRUCache) addNode(node *dLinkCache) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

// moveToHead 移动节点
func (l LRUCache) moveToHead(node *dLinkCache) {
	l.removeNode(node)
	l.addNode(node)
}

func (l LRUCache) removeTail() *dLinkCache {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
