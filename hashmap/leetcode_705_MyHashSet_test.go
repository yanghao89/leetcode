package hashmap

import "container/list"

type MyHashSet struct {
	data []list.List
}

func ConstructorSet() MyHashSet {
	return MyHashSet{
		data: make([]list.List, base),
	}
}
func (this *MyHashSet) Add(key int) {
	if !this.Contains(key) {
		h := this.hash(key)
		this.data[h].PushBack(key)
	}
}
func (this *MyHashSet) Remove(key int) {
	h := this.hash(key)
	for et := this.data[h].Front(); et != nil; et = et.Next() {
		if et.Value.(int) == key {
			this.data[h].Remove(et)
			return
		}
	}
}
func (this *MyHashSet) hash(key int) int {
	return key % base
}
func (this *MyHashSet) Contains(key int) bool {
	h := this.hash(key)
	for et := this.data[h].Front(); et != nil; et = et.Next() {
		if et.Value.(int) == key {
			return true
		}
	}
	return false
}
