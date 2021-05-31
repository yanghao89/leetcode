package leetcode_112

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(HasPathSum(root, 4))
}
