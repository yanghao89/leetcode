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
