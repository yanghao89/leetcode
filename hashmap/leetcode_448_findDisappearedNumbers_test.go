package hashmap

import (
	"fmt"
	"testing"
)

func findDisappearedNumbers01(nums []int) (ans []int) {
	n := len(nums)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v := range nums {
		if v <= n {
			ans = append(ans, i+1)
		}
	}
	return
}

func findDisappearedNumbers02(nums []int) (ans []int) {
	maps := make(map[int]struct{}, 0)
	for _, v := range nums {
		maps[v] = struct{}{}
	}
	for i := 1; i <= len(nums); i++ {
		if _, ok := maps[i]; !ok {
			ans = append(ans, i)
		}
	}
	return
}

func TestFindDisappearedNumbers(t *testing.T) {
	fmt.Println(findDisappearedNumbers01([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers02([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers02([]int{11, 12}))
}
