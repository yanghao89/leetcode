package array

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func findDuplicates(a []int) []int {
	sort.Ints(a)
	ints := []int{}
	for i := 1; i < len(a); i++ {
		if a[i-1] == a[i] {
			ints = append(ints, a[i])
		}
	}
	return ints
}

func findDuplicates02(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] != nums[nums[i]-1] {
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		} else {
			i++
		}
	}
	arr := []int{}
	for j := 0; j < len(nums); j++ {
		if nums[j] != j+1 {
			arr = append(arr, int(math.Abs(float64(nums[j]))))
		}
	}
	return arr
}

func TestFindDuplicates(t *testing.T) {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDuplicates02([]int{1, 1, 2}))
	fmt.Println(findDuplicates02([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
