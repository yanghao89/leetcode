package leetcode_145

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestPostorderTraversal(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	arr := PostorderTraversal(root)
	//a := true
	//b := false
	fmt.Println(arr)
}
