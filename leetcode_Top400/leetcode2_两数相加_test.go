package leetcode_Top400

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateListNode(nums []int) *ListNode {
	node := &ListNode{Val: nums[0], Next: nil}
	tmp := node
	for _, v := range nums[1:] {
		ne := &ListNode{Val: v}
		tmp.Next = ne
		tmp = ne
	}
	return node
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	carry := 0
	//拷贝内存地址
	list := head
	fmt.Printf("head:%p, list :%p", head, list)
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}
		list.Next = &ListNode{Val: carry % 10, Next: nil}
		carry /= 10
		list = list.Next
	}
	return head.Next
}

func TestAddTwoNumbers(t *testing.T) {
	data := addTwoNumbers(CreateListNode([]int{2, 4, 3}), CreateListNode([]int{5, 6, 4}))
	nums := make([]int, 0)
	for ; data != nil; data = data.Next {
		nums = append(nums, data.Val)
	}
	fmt.Println(nums)
}
