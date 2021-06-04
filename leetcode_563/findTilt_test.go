package leetcode_563

import (
	"fmt"
	"leetcode/utils"
	"testing"
)

func TestFindTilt(t *testing.T) {
	fmt.Println(FindTilt(utils.CreateTree([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})))
}
