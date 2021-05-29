package leetcode_101

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}

	left := 0
	right := len(arr) - 1
	//二分
	min := left + (right-left)>>1
	fmt.Println(arr[min])
	//node1 := utils.BuildTree([]int{1, 2, 3})
	//fmt.Println(IsSymmetric(node1))
}
