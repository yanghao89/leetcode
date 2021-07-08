package binary_search

import (
	"fmt"
	"testing"
)

func twoSum01(nums []int, target int) (ans []int) {

	for i := 0; i < len(nums); i++ {
		L, R := i+1, len(nums)-1
		for L <= R {
			mid := L + (R-L)>>1
			if nums[mid] == target-nums[i] {
				ans = []int{i + 1, mid + 1}
				return
			} else if nums[mid] > target-nums[i] {
				R = mid - 1
			} else {
				L = mid + 1
			}
		}
	}
	return []int{-1, -1}
}

func twoSum02(nums []int, target int) (ans []int) {
	L, R := 0, len(nums)-1
	for L < R {
		sum := nums[L] + nums[R]
		if sum == target {
			ans = []int{L + 1, R + 1}
			return
		} else if sum < target {
			L++
		} else {
			R--
		}
	}
	return []int{-1, -1}
}

func TestTwosum(t *testing.T) {
	fmt.Println(twoSum01([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum02([]int{2, 7, 11, 15}, 9))
}
