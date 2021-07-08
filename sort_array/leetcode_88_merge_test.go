package sort_array

import (
	"fmt"
	"sort"
	"testing"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(append(nums1[:m], nums2...))
}

type AvatarFrame struct {
	AvatarID  int64  `json:"avatar_id"`
	AvatarURI string `json:"avatar_uri"`
}

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	abc := new(AvatarFrame)
	//accc := 0
	fmt.Println(abc)
	//fmt.Println(nums1, nums2, unsafe.Sizeof(&abc.AvatarID), unsafe.Sizeof(accc))
}
