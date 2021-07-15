package medium

import (
	"fmt"
	"sort"
	"testing"
)

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	res := append(sort.IntSlice(nil), nums1...)
	res.Sort()
	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)
		sum += diff
		j := res.Search(v)
		if j < n {
			maxn = max(maxn, diff-(res[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-res[j-1]))
		}
	}
	return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func TestMinAbsoluteSumDiff(t *testing.T) {
	fmt.Println(minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5}))
	fmt.Println(minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4}))
}
