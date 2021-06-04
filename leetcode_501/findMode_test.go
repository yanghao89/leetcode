package leetcode_501

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestFindMode(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 5, 6, 7, 7, 8})
	arr := FindMode(root)
	fmt.Println(arr)
}
