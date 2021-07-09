package hashmap

import (
	"fmt"
	"math"
	"testing"
)

func majorityElement03(nums []int) (ans int) {
	ans = -1
	var (
		count int
	)
	for _, num := range nums {
		if count == 0 {
			ans = num
		}
		if ans == num {
			count++
		} else {
			count--
		}
	}
	count = 0
	for _, num := range nums {
		if num == ans {
			count++
		}
	}
	if count*2 < len(nums) {
		return -1
	}
	return
}

func majorityElement04(nums []int) (ans int) {
	maps := make(map[int]int, 0)
	for _, v := range nums {
		maps[v]++
	}
	max := math.MinInt32
	for k, _ := range maps {
		if max < maps[k] {
			max = maps[k]
			ans = k
		}
	}
	if max > len(nums)/2 {
		return
	}
	return -1
}

func TestMajorityElement01(t *testing.T) {
	fmt.Println(majorityElement03([]int{1, 2, 5, 9, 5, 9, 5, 5, 5}))
	fmt.Println(majorityElement03([]int{1, 2, 3}))
	fmt.Println(majorityElement04([]int{1, 2, 3}))
	fmt.Println(majorityElement04([]int{1, 2, 5, 9, 5, 9, 5, 5, 5}))
}
