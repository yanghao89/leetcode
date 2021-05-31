package leetcode_146_lru

import "container/list"

type LRUCaChe struct {
	size  int
	cache map[int]*list.Element
	list  *list.List
}

type Item struct {
	key, val int
}

func NewLRUCache(size int) *LRUCaChe {
	return &LRUCaChe{
		size:  size,
		cache: make(map[int]*list.Element),
		list:  list.New(),
	}
}

func (l *LRUCaChe) Get(key int) interface{} {
	if elem, ok := l.cache[key]; ok {
		//移动位置
		l.list.MoveToFront(elem)
		return elem.Value.(Item).val
	}
	return -1
}

func (l *LRUCaChe) Put(key, value int) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value = Item{key: key, val: value}
	} else {
		//
		if l.list.Len() >= l.size {
			element := l.list.Back()
			delete(l.cache, element.Value.(Item).key)
			l.list.Remove(element)
		} else {
			l.list.PushFront(Item{key: key, val: value})
			l.cache[key] = l.list.Front()
		}
	}
}
