package medium

import (
	"fmt"
	"testing"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	maps := map[*Node]*Node{}
	var dfs func(node *Node) *Node
	dfs = func(node *Node) *Node {
		if node == nil {
			return node
		}
		// 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
		if _, ok := maps[node]; ok {
			return maps[node]
		}
		// 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
		cloneNode := &Node{node.Val, []*Node{}}
		// 哈希表存储
		maps[node] = cloneNode
		// 遍历该节点的邻居并更新克隆节点的邻居列表
		for _, n := range node.Neighbors {
			cloneNode.Neighbors = append(cloneNode.Neighbors, dfs(n))
		}
		return cloneNode
	}
	return dfs(node)
}

func TestCloneGraph(t *testing.T) {
	node := &Node{}
	fmt.Println(cloneGraph(node))
}
