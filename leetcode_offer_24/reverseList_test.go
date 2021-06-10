package leetcode_offer_24

import (
	"leetcode/utils"
	"testing"
)

func BenchmarkReverseList(b *testing.B) {
	arrs := make([][]int, 5)
	arrs[0] = []int{1, 2, 3, 4, 5, 6, 7}
	arrs[1] = []int{8, 9, 10, 11, 12, 13, 14}
	arrs[2] = []int{15, 16, 17, 18, 19, 20, 21}
	arrs[3] = []int{22, 23, 24, 25, 26, 27, 28}
	arrs[4] = []int{29, 30, 31, 32, 33, 34, 35}
	//var n int = 0
	//for i := 0; i < b.N; i++ {
	//	fmt.Println(n)
	//	//if len(arrs[n]) > 0 {
	//	//	node := utils.CreateListNode(arrs[n])
	//	//	//ReverseList(node)
	//	//	fmt.Println(node)
	//	//}
	//	n++
	//}
	num := 4
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if num == -1 {
			break
		}
		node := utils.CreateListNode(arrs[num])
		ReverseList(node)
		//fmt.Sprintf("%d", num)
		num--
	}
	//node := utils.CreateListNode([]int{1, 2, 3, 4, 5, 6, 7})
	//fmt.Println(node)
	//
	//fmt.Println(ReverseList(node))
}

func ReverseList(head *utils.ListNode) *utils.ListNode {
	var pre *utils.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
