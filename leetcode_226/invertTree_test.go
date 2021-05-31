package leetcode_226

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestInvertTree(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(root)
	node := InvertTree(root)
	fmt.Println(node)
}
