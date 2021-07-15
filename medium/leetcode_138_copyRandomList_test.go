package medium

type RandomNode struct {
	Val    int
	Next   *RandomNode
	Random *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
	h := make(map[*RandomNode]*RandomNode, 0)
	for p := head; p != nil; p = p.Next {
		h[p] = &RandomNode{Val: p.Val, Next: nil, Random: nil}
	}
	for p := head; p != nil; p = p.Next {
		h[p].Next = h[p.Next]
		h[p].Random = h[p.Random]
	}
	return h[head]
}
