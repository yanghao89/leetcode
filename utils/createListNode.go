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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	i = int(-1)
)

func BuildTree(arr []int) *TreeNode {
	i += 1
	if i >= len(arr) {
		return nil
	}
	t := new(TreeNode)
	if arr[i] != 0 {
		t.Val = arr[i]
		t.Left = BuildTree(arr)
		t.Right = BuildTree(arr)
	} else {
		return nil
	}
	return t
}
