package array

import (
	"fmt"
	"testing"
)

func smallerNumbersThanCurrent(nums []int) (ans []int) {
	for _, v := range nums {
		cnt := 0
		for _, vv := range nums {
			if vv < v {
				cnt++
			}
		}
		ans = append(ans, cnt)
	}
	return
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}
