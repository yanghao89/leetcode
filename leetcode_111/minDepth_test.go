package leetcode_111

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestMinDepth(t *testing.T) {
	root := utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 9})
	fmt.Println(MinDepth(root))
}
