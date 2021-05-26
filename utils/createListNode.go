package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(arr []int) *ListNode {
	node := &ListNode{Val: 0}
	tmp := node
	for i := 0; i < len(arr); i++ {
		ne := &ListNode{Val: arr[i]}
		tmp.Next = ne
		tmp = ne
	}
	return node.Next
}
