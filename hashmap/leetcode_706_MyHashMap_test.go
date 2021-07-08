package hashmap

import "container/list"

var (
	base = 1024
)

type MyHashMap struct {
	data []list.List
}

type entry struct {
	key, value int
}

func Constructor() MyHashMap {
	return MyHashMap{
		data: make([]list.List, base),
	}
}
func (this *MyHashMap) hash(key int) int {
	return key % base
}

func (this *MyHashMap) Put(key int, value int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			e.Value = entry{key: key, value: value}
			return
		}
	}
	this.data[h].PushBack(entry{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			return et.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	h := this.hash(key)
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(entry); et.key == key {
			this.data[h].Remove(e)
		}
	}
}
