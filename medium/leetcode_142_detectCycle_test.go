package medium

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	maps := make(map[*ListNode]struct{}, 0)
	for head != nil {
		if _, ok := maps[head]; ok {
			return head
		}
		maps[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle01(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

func TestDetectCycle(t *testing.T) {
	fmt.Println(detectCycle(&ListNode{}))
}
