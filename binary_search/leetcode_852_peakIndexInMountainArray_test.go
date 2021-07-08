package binary_search

import (
	"fmt"
	"testing"
)

func peakIndexInMountainArray(arr []int) (ans int) {
	L, R := 0, len(arr)-1
	for L <= R {
		min := L + (R-L)>>1
		if arr[min] > arr[min+1] {
			ans = min
			R = min - 1
		} else {
			L = min + 1
		}
	}
	return
}

func TestPeakIndexInMountainArray(t *testing.T) {
	fmt.Println(peakIndexInMountainArray([]int{0, 1, 0}))
}
