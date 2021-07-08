package sort_array

import (
	"fmt"
	"testing"
)

func QuickSort(num []int) []int {
	qSort(num, 0, len(num)-1)
	return num
}

func qSort(nums []int, L, R int) {
	if L > R {
		return
	}
	i, j, base := L, R, nums[L]
	for i < j {
		for nums[j] >= base && i < j {
			j--
		}
		for nums[i] <= base && i < j {
			i++
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[i], nums[L] = nums[L], nums[i]
	qSort(nums, L, i-1)
	qSort(nums, i+1, R)
}

func TestQuickSort(t *testing.T) {
	fmt.Println(QuickSort([]int{5, 2, 3, 1}))
	//var (
	//	a string
	//)
	//fmt.Println(unsafe.Sizeof(a))

}
