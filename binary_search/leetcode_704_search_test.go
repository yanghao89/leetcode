package binary_search

import (
	"fmt"
	"sync"
	"testing"
)

func search(nums []int, target int) (ans int) {
	ans = -1
	L, R := 0, len(nums)-1
	for L <= R {
		min := L + (R-L)>>1
		if nums[min] == target {
			return min
		} else if nums[min] < target {
			L = min + 1
		} else {
			R = min - 1
		}
	}
	return ans
}

func TestSearch(t *testing.T) {
	var (
		wg sync.WaitGroup
	)
	wg.Add(2)
	ch := make(chan int)
	go func() {
		wg.Done()
		nums := make([][]int, 0)
		nums = append(nums, []int(nil), []int(nil), []int(nil))
		num1 := []int{-1, 0, 3, 5, 9, 12}
		nums[0] = append(nums[0], num1...)
		nums[1] = append(nums[1], num1...)
		nums[2] = append(nums[2], []int{5}...)
		targets := []int{9, 2, 5}
		for i := 0; i < len(nums); i++ {
			ch <- search(nums[i], targets[i])
		}
		close(ch)
	}()
	go func() {
		wg.Done()
		for {
			select {
			case vv, ok := <-ch:
				if !ok {
					break
				}
				fmt.Println(vv)
			}
		}
	}()
	wg.Wait()
}
