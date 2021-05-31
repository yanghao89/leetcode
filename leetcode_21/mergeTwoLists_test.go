package leetcode_21

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	res := MergeTwoLists(utils.CreateListNode([]int{1, 2, 2, 3, 3, 4, 5}), utils.CreateListNode([]int{6, 7, 8, 9, 10}))
	next := res.Next
	fmt.Println(res.Val)
	for next.Next != nil {
		next = next.Next
		fmt.Println(next.Val)
	}
}
