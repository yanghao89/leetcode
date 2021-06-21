package deleteDuplicates

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	node := utils.CreateListNode([]int{1, 1, 2, 3, 3})
	//for _,v := range node{
	//	fmt.Println(v)
	//}
	data := DeleteDuplicates(node)
	for data != nil {
		fmt.Println(data.Val)
		data = data.Next
	}
	//
	fmt.Println(data)
}

func TestDelete(t *testing.T) {
	arr := Delete([]int{1, 1, 2, 2, 3, 3, 4, 5, 6, 6})
	fmt.Println(arr)
}

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
