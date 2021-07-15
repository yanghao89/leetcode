package hashmap

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func getIntersectionNode01(headA, headB *utils.ListNode) *utils.ListNode {
	maps := make(map[*utils.ListNode]struct{}, 0)
	for ; headA != nil; headA = headA.Next {
		maps[headA] = struct{}{}
	}
	for ; headB != nil; headB = headB.Next {
		if _, ok := maps[headB]; ok {
			return headB
		}
	}
	return nil
}

func TestGetIntersectionNode01(t *testing.T) {
	fmt.Println(getIntersectionNode01(utils.CreateListNode([]int{4, 1, 8, 4, 5}, false), utils.CreateListNode([]int{5, 0, 1, 8, 4, 5}, false)))
}
