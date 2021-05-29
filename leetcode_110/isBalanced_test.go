package leetcode_110

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(IsBalanced(root))

}
