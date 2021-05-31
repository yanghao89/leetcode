package leetcode_144

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	arr := PreorderTraversal(root)
	//a := true
	//b := false
	fmt.Println(arr)
}
