package leetcode_606

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestTree2str(t *testing.T) {
	fmt.Println(Tree2str(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}
