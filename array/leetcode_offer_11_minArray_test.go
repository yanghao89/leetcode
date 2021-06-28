package array

import (
	"fmt"
	"testing"
)

func TestMinArray(t *testing.T) {
	fmt.Println(MinArray([]int{3, 4, 5, 1, 2}))
}

func MinArray(numbers []int) int {
	L, R := 0, len(numbers)-1
	for L <= R {
		mid := L + (R-L)>>1
		//如果mid 小于右边 , 指针向左
		if numbers[mid] < numbers[R] {
			R = mid
			//如果mid大于左边 , 指针继续向左
		} else if numbers[mid] > numbers[R] {
			L = mid + 1
		} else {
			R -= 1
		}
	}
	return numbers[L]
}
