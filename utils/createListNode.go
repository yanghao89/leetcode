package utils

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(arr []int) *ListNode {
	node := &ListNode{Val: arr[0]}
	tmp := node
	for i := 0; i < len(arr[1:])+1; i++ {
		ne := &ListNode{Val: arr[i]}
		tmp.Next = ne
		tmp = ne
	}
	return node
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTree(arr []int) *TreeNode {
	return create(arr, 0, len(arr)-1)
}

func create(arr []int, L, R int) *TreeNode {
	if L > R {
		return nil
	}
	mid := L + (R-L)>>1
	root := &TreeNode{Val: arr[mid]}
	//右边向左一位
	root.Left = create(arr, L, mid-1)
	//左边想向后一位
	root.Right = create(arr, mid+1, R)
	return root
}
