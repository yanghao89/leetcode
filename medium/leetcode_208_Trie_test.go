package medium

import "testing"

/**
Trie，又称前缀树或字典树，是一棵有根树
*/
type Trie struct {
	isEnd    bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{
		isEnd:    false,
		children: [26]*Trie{},
	}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		ch -= 'a'
		if node.children[ch] == nil {
			node.children[ch] = &Trie{}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}
func (this *Trie) searchPrefix(prefix string) *Trie {
	node := this
	for _, ch := range prefix {
		ch -= 'a'
		if node.children[ch] == nil {
			return nil
		}
		node = node.children[ch]
	}
	return node
}

func (this *Trie) Search(word string) bool {
	node := this.searchPrefix(word)
	return node != nil && node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func TestTrie(t *testing.T) {
	Trie := Constructor()
	Trie.Insert("apple")
	Trie.Search("apple")
}
