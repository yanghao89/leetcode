package deleteDuplicates

import (
	"leetcode/utils"
)

func DeleteDuplicates(head *utils.ListNode) *utils.ListNode {
	//判断 链表节点是否为nil
	if head == nil {
		return head
	}
	//copy 链表
	currNode := head
	//判断当前节点下一个指针是否为空
	for currNode.Next != nil {
		// 当前节点的val 和当前节点下一个指针的val 是否相等 如果相等,
		//当前节点下一个指针的下一个指针交换到当前节点下一个指针
		if currNode.Val == currNode.Next.Val {
			currNode.Next = currNode.Next.Next
		} else {
			currNode = currNode.Next
		}
	}
	return head
}

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

func Delete(arr []int) []int {
	news, n := []int{arr[0]}, 1
	for i := 0; i < len(arr); i++ {
		if n > len(arr)-1 {
			break
		}
		if arr[i] != arr[n] {
			news = append(news, arr[n])
		}
		n++
	}
	return news
}
