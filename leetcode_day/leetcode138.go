package day

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	m := make(map[*Node]*Node, 0)
	for p := head; p != nil; p = p.Next {
		m[p] = &Node{Val: p.Val, Next: nil, Random: nil}
	}
	for p := head; p != nil; p = p.Next {
		m[p].Next = m[p.Next]
		m[p].Random = m[p.Random]
	}
	return m[head]
}
